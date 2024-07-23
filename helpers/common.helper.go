package helpers

import (
	"github.com/joho/godotenv"
	"os"
)

func GetEnv(envKey string, defaultValue string) string {
	godotenv.Load()

	val := os.Getenv(envKey)
	if val == "" {
		val = defaultValue
	}

	return val
}
