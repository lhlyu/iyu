package middleware

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/errcode"
)

func Permission() context.Handler {
	return func(ctx iris.Context) {
		user := ctx.Values().Get(common.X_USER).(*common.XUser)
		if user.Role < common.PERMISSION {
			ctx.JSON(errcode.NoPermission)
			return
		}
		ctx.Values().Set(common.ADMIN, true)
		ctx.Next()
	}
}
