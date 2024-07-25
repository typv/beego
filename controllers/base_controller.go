package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"net/http"
	"src/ultils"
)

type BaseController struct {
	web.Controller
}

func (bc *BaseController) ResponseOk(data any) {
	statusCode := http.StatusOK
	bc.Ctx.Output.SetStatus(statusCode)
	bc.Ctx.Output.Header("Content-Type", "application/json")
	response := ultils.Response{
		Code: statusCode,
		Data: data,
	}
	bc.Data["json"] = response
	bc.ServeJSON()
}

func (bc *BaseController) ResponseErr(statusCode int, message any, data any) {
	bc.Ctx.Output.SetStatus(statusCode)
	bc.Ctx.Output.Header("Content-Type", "application/json")
	response := ultils.ErrorResponse{
		Code:    statusCode,
		Message: message,
		Data:    data,
	}

	bc.Data["json"] = response
	bc.ServeJSON()
}

func (bc *BaseController) ResponseUnAuthorized(message any, data any) {
	statusCode := http.StatusUnauthorized
	if message == nil {
		message = "Unauthorized."
	}
	bc.ResponseErr(statusCode, message, data)
}

func (bc *BaseController) ResponseNotFound(message any, data any) {
	statusCode := http.StatusNotFound
	if message == nil {
		message = "Object not found."
	}
	bc.ResponseErr(statusCode, message, data)
}

func (bc *BaseController) ResponseBadRequest(message any, data any) {
	statusCode := http.StatusBadRequest
	if message == nil {
		message = "Bad request."
	}
	bc.ResponseErr(statusCode, message, data)
}

func (bc *BaseController) ResponseForbidden(message any, data any) {
	statusCode := http.StatusForbidden
	if message == nil {
		message = "Forbidden."
	}
	bc.ResponseErr(statusCode, message, data)
}

func (bc *BaseController) ResponseInternalServerError(message any, data any) {
	statusCode := http.StatusInternalServerError
	if message == nil {
		message = "Internal server error."
	}
	bc.ResponseErr(statusCode, message, data)
}
