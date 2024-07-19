package routers

import (
	"github.com/beego/beego/v2/server/web"
	"src/controllers"
)

func init() {
	web.Router("/", &controllers.MainController{})

	ns := web.NewNamespace("/v1",
		web.NSNamespace("/user",
			web.NSRouter("/", &controllers.UserController{}, "get:GetHello"),
		),
	)

	web.AddNamespace(ns)
}
