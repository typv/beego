package helpers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("helpers.GetEnv: Error loading .env file", err)
	}
}

func GetEnv(envKey string, defaultValue string) string {
	if value, exists := os.LookupEnv(envKey); exists {
		return value
	}
	return defaultValue
}
