package helpers

import (
	"github.com/joho/godotenv"
	"os"
)

func GetEnv(envKey string, defaultValue string) string {
	godotenv.Load()
	if value, exists := os.LookupEnv(envKey); exists {
		return value
	}
	return defaultValue
}
