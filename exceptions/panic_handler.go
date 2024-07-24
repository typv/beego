package exceptions

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
)

func init() {
	web.BConfig.RecoverPanic = true
	web.BConfig.RecoverFunc = func(context *context.Context, config *web.Config) {
		if err := recover(); err != nil {
			errCode := http.StatusInternalServerError
			errMsg := fmt.Sprintf("ERR: %v", err)
			response := map[string]interface{}{
				"code":    errCode,
				"message": errMsg,
			}
			context.Output.Header("Content-Type", "application/json")
			context.Output.SetStatus(errCode)
			context.Output.JSON(response, false, false)
		}
	}
}
