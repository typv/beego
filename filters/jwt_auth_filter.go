package filters

import (
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
	"src/exceptions"
	"src/services"
	"strings"
)

func JWTAuthFilter(ctx *context.Context) {
	if skipAuth, ok := ctx.Input.GetData("public").(bool); ok && skipAuth {
		return
	}

	authHeader := ctx.Input.Header("Authorization")
	if authHeader == "" {
		panic(exceptions.PanicError{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	JWTHelper := services.NewJwtService()
	token, err := JWTHelper.Sign(tokenString)

	if err != nil || !token.Valid {
		panic(exceptions.PanicError{
			Message: "Unauthorized",
			Code:    http.StatusUnauthorized,
		})
	}
}
