package routers

import (
	"github.com/beego/beego/v2/server/web"
	"src/app/controllers"
	"src/filters"
)

func init() {
	web.Router("/", &controllers.MainController{})

	ns := web.NewNamespace("/v1",
		web.NSNamespace("/users",
			web.NSRouter("/", &controllers.UserController{}, "get:GetUsers"),
			web.NSRouter("/detail", &controllers.UserController{}, "get:GetUser"),
		),
	)

	web.AddNamespace(ns)

	// Public API
	web.InsertFilter("/", web.BeforeRouter, filters.PublicFilter) // Home route

	// Apply JWT filter
	web.InsertFilter("*", web.BeforeRouter, filters.JWTAuthFilter)
}
