package config

import (
	"fmt"
	"src/helpers"
)

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	TimeZone string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     helpers.GetEnv("DB_HOST", "localhost"),
		User:     helpers.GetEnv("DB_USERNAME", ""),
		Password: helpers.GetEnv("DB_PASSWORD", ""),
		DBName:   helpers.GetEnv("DB_DATABASE", ""),
		Port:     helpers.GetEnv("DB_PORT", "5432"),
		SSLMode:  helpers.GetEnv("DB_SSLMODE", ""),
		TimeZone: helpers.GetEnv("DB_TIMEZONE", "UTC"),
	}
}

func DBConnectString() string {
	config := NewDatabaseConfig()
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode, config.TimeZone)
}
