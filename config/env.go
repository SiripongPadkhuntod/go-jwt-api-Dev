package config

import (
	"log"

	"github.com/joho/godotenv"

	"os"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		if os.Getenv("APP_ENV") == "dev" {
			log.Println("⚠️ No .env file found")
		}
	}
}