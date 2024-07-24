package filters

import (
	"github.com/beego/beego/v2/server/web/context"
)

func PublicFilter(ctx *context.Context) {
	ctx.Input.SetData("public", true)
	return
}
