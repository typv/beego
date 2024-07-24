package main

import (
	"github.com/beego/beego/v2/server/web"
	_ "github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	_ "src/database"
	_ "src/exceptions"
	"src/helpers"
	_ "src/routers"
)

func main() {
	port := helpers.GetEnv("APP_PORT", "8080")

	web.Run(":" + port)
}
