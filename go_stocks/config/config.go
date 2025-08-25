package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	ExternalAPI string
	Port        string
	Key         string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file")
	}

	return Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		ExternalAPI: os.Getenv("API_URL"),
		Port:        os.Getenv("PORT"),
		Key:         os.Getenv("API_KEY"),
	}
}
