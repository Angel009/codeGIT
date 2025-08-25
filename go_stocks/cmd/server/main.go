package main

import (
	"log"
	"stocks/config"
	"stocks/internal/db"
	"stocks/internal/router"
)

func main() {
	cfg := config.LoadConfig()

	database, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Error conecting to DB:", err)
	}
	defer database.Close()

	r := router.Setup(database, cfg.ExternalAPI)

	log.Println("---http://localhost:8080---")
	if err := r.Run(cfg.Port); err != nil {
		log.Fatal(err)
	}
}
