package config

import (
	"src/helpers"
)

type JWTConfig struct {
	Secret  string
	Expired string
}

func NewJWTConfig() *JWTConfig {
	return &JWTConfig{
		Secret:  helpers.GetEnv("JWT_SECRET", ""),
		Expired: helpers.GetEnv("JWT_EXPIRED_MINUTES", "720"),
	}
}
