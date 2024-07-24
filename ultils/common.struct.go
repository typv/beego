package ultils

import "github.com/golang-jwt/jwt/v5"

type AuthUser struct {
	ID    string `json:"sub"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
	jwt.RegisteredClaims
}

type Response struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

type ErrorResponse struct {
	Code    int `json:"code"`
	Message any `json:"message"`
	Data    any `json:"data,omitempty"`
}

type PanicError struct {
	Message string
	Code    any
}
