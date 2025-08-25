package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"stocks/internal/repository"
	"stocks/internal/services"

	"github.com/gin-gonic/gin"
)

func GetApiData(apiURL string, apiDataRepo *repository.ApiDataRepository, analysisService *services.AnalysisService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Hacer request a la API externa
		req, err := http.NewRequest("GET", apiURL, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating request"})
			return
		}

		apiKey := os.Getenv("API_KEY")
		req.Header.Set("Authorization", "Bearer "+apiKey)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error calling API"})
			return
		}
		defer resp.Body.Close()

		var data interface{}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error decoding data"})
			return
		}

		if err := apiDataRepo.UpsertApiData(data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error saving in database"})
			return
		}

		savedData, err := apiDataRepo.GetLatestApiData()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting data"})
			return
		}

		var responseData interface{}
		if err := json.Unmarshal([]byte(savedData.Data), &responseData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error processing data"})
			return
		}

		dataMap, ok := responseData.(map[string]interface{})
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid data format"})
			return
		}

		analysis, err := analysisService.FindBestStock(dataMap)
		if err != nil {
			// If FindBestStock fails, send only stocks data
			c.JSON(http.StatusOK, gin.H{
				"data":           responseData,
				"updated_at":     savedData.UpdatedAt,
				"analysis_error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":       responseData,
			"best":       analysis,
			"updated_at": savedData.UpdatedAt,
		})
	}
}
