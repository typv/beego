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

func (this *BaseController) ResponseErr(code int, message string) {
	this.Ctx.Output.SetStatus(code)
	this.Ctx.Output.Header("Content-Type", "application/json")
	response := ErrorResponse{
		Code:    code,
		Message: message,
	}
	this.Data["json"] = response
	this.ServeJSON()
}

func (this *BaseController) ResponseOk(data any) {
	statusCode := http.StatusOK
	this.Ctx.Output.SetStatus(statusCode)
	this.Ctx.Output.Header("Content-Type", "application/json")
	response := Response{
		Code: statusCode,
		Data: data,
	}
	this.Data["json"] = response
	this.ServeJSON()
}
