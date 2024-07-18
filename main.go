package main

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
	"log"
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
	fmt.Println(portInt)
	beego.BConfig.Listen.HTTPPort = portInt

	beego.Run()
}
