package exceptions

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
	"src/ultils"
)

func init() {
	web.BConfig.RecoverPanic = true
	web.BConfig.RecoverFunc = func(context *context.Context, config *web.Config) {
		if err := recover(); err != nil {
			var errMsg string
			var errCode int

			switch e := err.(type) {
			case string:
				errMsg = e
			case error:
				errMsg = e.Error()
			case ultils.PanicError:
				errMsg = e.Message
				errCode, _ = e.Code.(int)
			default:
				errMsg = "unknown error"
			}

			if errCode == 0 {
				errCode = http.StatusInternalServerError
			}

			response := ultils.ErrorResponse{
				Code:    errCode,
				Message: errMsg,
			}
			context.Output.Header("Content-Type", "application/json")
			context.Output.SetStatus(errCode)
			context.Output.JSON(response, false, false)
		}
	}
}
