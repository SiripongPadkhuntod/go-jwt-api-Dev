package config

import (
	"log"
	"os"
)

func GetJwtSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is not set")
	}
	return []byte(secret)
}
