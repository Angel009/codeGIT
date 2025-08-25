package models

import "time"

type ApiData struct {
	ID        int       `json:"id" db:"id"`
	Data      string    `json:"data" db:"data"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
