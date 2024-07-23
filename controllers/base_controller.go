package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BaseController struct {
	web.Controller
}

func (c *BaseController) ResponseErr(code int, message string) {
	c.Ctx.Output.SetStatus(code)
	c.Ctx.Output.Header("Content-Type", "application/json")
	response := ErrorResponse{
		Code:    code,
		Message: message,
	}
	c.Data["json"] = response
	c.ServeJSON()
}

func (c *BaseController) ResponseOk(data any) {
	statusCode := http.StatusOK
	c.Ctx.Output.SetStatus(statusCode)
	c.Ctx.Output.Header("Content-Type", "application/json")
	response := Response{
		Code: statusCode,
		Data: data,
	}
	c.Data["json"] = response
	c.ServeJSON()
}
