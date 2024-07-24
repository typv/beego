package main

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	_ "src/database"
	_ "src/exceptions"
	"src/helpers"
	_ "src/routers"
	"strconv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := helpers.GetEnv("APP_PORT", "8080")
	portInt, _ := strconv.Atoi(port)
	web.BConfig.Listen.HTTPPort = portInt

	web.Run()
}
