package services

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type StockItem struct {
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	Company    string `json:"company"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Ticker     string `json:"ticker"`
	Time       string `json:"time"`
}

type AnalysisResult struct {
	BestItem        *StockItem `json:"best_item"`
	Criteria        string     `json:"criteria"`
	TargetIncrease  float64    `json:"target_increase"`
	PercentIncrease float64    `json:"percent_increase"`
	TotalItems      int        `json:"total_items"`
}

type AnalysisService struct{}

func NewAnalysisService() *AnalysisService {
	return &AnalysisService{}
}

// Analize stocks and return best stock
func (s *AnalysisService) FindBestStock(data map[string]interface{}) (*AnalysisResult, error) {

	// Extract from map
	itemsInterface, exists := data["items"]
	if !exists {
		return nil, fmt.Errorf("there's no stocks here")
	}

	itemsSlice, ok := itemsInterface.([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid array")
	}

	//Store as structs
	var stockItems []StockItem
	for _, item := range itemsSlice {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		stockItem := StockItem{}
		if val, ok := itemMap["action"].(string); ok {
			stockItem.Action = val
		}
		if val, ok := itemMap["brokerage"].(string); ok {
			stockItem.Brokerage = val
		}
		if val, ok := itemMap["company"].(string); ok {
			stockItem.Company = val
		}
		if val, ok := itemMap["rating_from"].(string); ok {
			stockItem.RatingFrom = val
		}
		if val, ok := itemMap["rating_to"].(string); ok {
			stockItem.RatingTo = val
		}
		if val, ok := itemMap["target_from"].(string); ok {
			stockItem.TargetFrom = val
		}
		if val, ok := itemMap["target_to"].(string); ok {
			stockItem.TargetTo = val
		}
		if val, ok := itemMap["ticker"].(string); ok {
			stockItem.Ticker = val
		}
		if val, ok := itemMap["time"].(string); ok {
			stockItem.Time = val
		}

		stockItems = append(stockItems, stockItem)
	}

	if len(stockItems) == 0 {
		return nil, fmt.Errorf("stocks not found")
	}

	return s.analyzeBestStock(stockItems)
}

func (s *AnalysisService) analyzeBestStock(items []StockItem) (*AnalysisResult, error) {
	var bestItem *StockItem
	var maxIncrease float64
	var maxPercentIncrease float64

	for i := range items {
		item := &items[i]

		targetFrom := s.parsePrice(item.TargetFrom)
		targetTo := s.parsePrice(item.TargetTo)

		if targetFrom > 0 && targetTo > 0 {
			increase := targetTo - targetFrom
			percentIncrease := (increase / targetFrom) * 100

			isBetter := false

			if increase > maxIncrease {
				isBetter = true
			} else if increase == maxIncrease && bestItem != nil {

				if s.isMoreRecent(item.Time, bestItem.Time) {
					isBetter = true
				} else if item.Time == bestItem.Time {

					if s.isBetterRating(item.RatingTo, bestItem.RatingTo) {
						isBetter = true
					}
				}
			}

			if isBetter {
				bestItem = item
				maxIncrease = increase
				maxPercentIncrease = percentIncrease
			}
		}
	}

	if bestItem == nil {
		bestItem = s.findMostRecentWithBestRating(items)
		if bestItem != nil {
			targetFrom := s.parsePrice(bestItem.TargetFrom)
			targetTo := s.parsePrice(bestItem.TargetTo)
			if targetFrom > 0 && targetTo > 0 {
				maxIncrease = targetTo - targetFrom
				maxPercentIncrease = (maxIncrease / targetFrom) * 100
			}
		}
	}

	if bestItem == nil {
		return nil, fmt.Errorf("unable to determine best stock")
	}

	return &AnalysisResult{
		BestItem:        bestItem,
		Criteria:        "Highest target increase + recent date + rating",
		TargetIncrease:  maxIncrease,
		PercentIncrease: maxPercentIncrease,
		TotalItems:      len(items),
	}, nil
}

// Converts string to float64
func (s *AnalysisService) parsePrice(priceStr string) float64 {
	cleaned := strings.ReplaceAll(strings.TrimSpace(priceStr), "$", "")
	price, err := strconv.ParseFloat(cleaned, 64)
	if err != nil {
		return 0
	}
	return price
}

// Compares timestamps
func (s *AnalysisService) isMoreRecent(time1, time2 string) bool {
	t1, err1 := time.Parse(time.RFC3339, time1)
	t2, err2 := time.Parse(time.RFC3339, time2)

	if err1 != nil || err2 != nil {
		return false
	}

	return t1.After(t2)
}

// Compares stock ratings
func (s *AnalysisService) isBetterRating(rating1, rating2 string) bool {
	ratingScore := map[string]int{
		"Outperform":   5,
		"Buy":          4,
		"Equal Weight": 3,
		"Neutral":      2,
		"Underweight":  1,
	}

	score1, exists1 := ratingScore[rating1]
	score2, exists2 := ratingScore[rating2]

	if !exists1 || !exists2 {
		return false
	}

	return score1 > score2
}

// Focus on most recent stock with best rating
func (s *AnalysisService) findMostRecentWithBestRating(items []StockItem) *StockItem {
	var best *StockItem
	var bestTime time.Time

	for i := range items {
		item := &items[i]

		itemTime, err := time.Parse(time.RFC3339, item.Time)
		if err != nil {
			continue
		}

		if best == nil || itemTime.After(bestTime) ||
			(itemTime.Equal(bestTime) && s.isBetterRating(item.RatingTo, best.RatingTo)) {
			best = item
			bestTime = itemTime
		}
	}

	return best
}
