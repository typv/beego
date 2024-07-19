package controllers

import "github.com/beego/beego/v2/server/web"

type UserController struct {
	web.Controller
}

func (c *UserController) GetHello() {
	response := map[string]string{"message": "Hello, User!"}
	c.Data["json"] = response
	c.ServeJSON()
}
