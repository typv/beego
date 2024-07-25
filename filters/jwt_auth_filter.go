package filters

import (
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
	"src/app/services"
	"src/ultils"
	"strings"
)

func JWTAuthFilter(ctx *context.Context) {
	if skipAuth, ok := ctx.Input.GetData("public").(bool); ok && skipAuth {
		return
	}

	authHeader := ctx.Input.Header("Authorization")
	if authHeader == "" {
		panic(ultils.PanicError{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	JwtService := services.NewJwtService()
	token, err, authUser := JwtService.Sign(tokenString)

	if err != nil || !token.Valid {
		panic(ultils.PanicError{
			Message: "Unauthorized",
			Code:    http.StatusUnauthorized,
		})
	}
	ctx.Input.SetData("authUser", authUser)
}
