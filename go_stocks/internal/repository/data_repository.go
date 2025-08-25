package repository

import (
	"database/sql"
	"encoding/json"
	"stocks/internal/models"
	"time"
)

type ApiDataRepository struct {
	DB *sql.DB
}

func NewApiDataRepository(db *sql.DB) *ApiDataRepository {
	return &ApiDataRepository{DB: db}
}

func (r *ApiDataRepository) UpsertApiData(data interface{}) error {
	// Convert map[string]interface{} to JSON string
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	now := time.Now()

	query := `
		INSERT INTO data_api (id, data, created_at, updated_at)
		VALUES (1, $1, $2, $3)
		ON CONFLICT (id) 
		DO UPDATE SET 
			data = EXCLUDED.data,
			updated_at = EXCLUDED.updated_at
	`

	_, err = r.DB.Exec(query, string(jsonData), now, now)
	return err
}

func (r *ApiDataRepository) GetLatestApiData() (*models.ApiData, error) {
	var apiData models.ApiData
	query := `
		SELECT id, data, created_at, updated_at 
		FROM data_api 
		WHERE id = 1
	`

	err := r.DB.QueryRow(query).Scan(
		&apiData.ID,
		&apiData.Data,
		&apiData.CreatedAt,
		&apiData.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &apiData, nil
}
