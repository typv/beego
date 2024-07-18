package helpers

import "os"

func GetEnv(envKey string, defaultValue string) string {
	val := os.Getenv("APP_PORT")
	if val == "" {
		val = defaultValue
	}

	return val
}
