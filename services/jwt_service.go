package services

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"src/config"
	"strconv"
	"time"
)

type IJwtService interface {
	Sign(token string) (*jwt.Token, error)
	Claims(data map[string]interface{}) (string, error)
}

type JwtService struct{}

var (
	jwtSecret []byte
	jwtExpiry time.Duration
)

func init() {
	jwtSecret = []byte(config.NewJWTConfig().Secret)
	expiryStr := config.NewJWTConfig().Expired
	expiryInt, _ := strconv.Atoi(expiryStr)
	jwtExpiry = time.Duration(expiryInt) * time.Minute
}

func NewJwtService() IJwtService {
	return &JwtService{}
}

func (this *JwtService) Sign(token string) (*jwt.Token, error) {
	singedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrAbortHandler
		}
		return jwtSecret, nil
	})

	return singedToken, err
}

func (this *JwtService) Claims(data map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{}
	for key, value := range data {
		claims[key] = value
	}
	claims["exp"] = time.Now().Add(jwtExpiry).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, err
}
