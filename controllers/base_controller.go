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
	Code    int `json:"code"`
	Message any `json:"message"`
	Data    any `json:"data,omitempty"`
}

type BaseController struct {
	web.Controller
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

func (this *BaseController) ResponseErr(statusCode int, message any, data any) {
	this.Ctx.Output.SetStatus(statusCode)
	this.Ctx.Output.Header("Content-Type", "application/json")
	response := ErrorResponse{
		Code:    statusCode,
		Message: message,
		Data:    data,
	}

	this.Data["json"] = response
	this.ServeJSON()
}

func (this *BaseController) ResponseUnAuthorized(message any, data any) {
	statusCode := http.StatusUnauthorized
	if message == nil {
		message = "Unauthorized."
	}
	this.ResponseErr(statusCode, message, data)
}

func (this *BaseController) ResponseNotFound(message any, data any) {
	statusCode := http.StatusNotFound
	if message == nil {
		message = "Object not found."
	}
	this.ResponseErr(statusCode, message, data)
}

func (this *BaseController) ResponseBadRequest(message any, data any) {
	statusCode := http.StatusBadRequest
	if message == nil {
		message = "Bad request."
	}
	this.ResponseErr(statusCode, message, data)
}

func (this *BaseController) ResponseForbidden(message any, data any) {
	statusCode := http.StatusForbidden
	if message == nil {
		message = "Forbidden."
	}
	this.ResponseErr(statusCode, message, data)
}

func (this *BaseController) ResponseInternalServerError(message any, data any) {
	statusCode := http.StatusInternalServerError
	if message == nil {
		message = "Internal server error."
	}
	this.ResponseErr(statusCode, message, data)
}
