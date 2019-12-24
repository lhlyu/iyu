package middleware

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/errcode"
)

// for login user
func PermissionUser() context.Handler {
	return func(ctx iris.Context) {
		user, ok := ctx.Values().Get(common.X_USER).(*common.XUser)
		if !ok || user == nil || user.Id == 0 {
			ctx.JSON(errcode.NoLogin)
			return
		}
		ctx.Next()
	}
}

// for admin
func PermissionAdmin() context.Handler {
	return func(ctx iris.Context) {
		user, ok := ctx.Values().Get(common.X_USER).(*common.XUser)
		if !ok || user == nil || user.Id == 0 {
			ctx.JSON(errcode.NoLogin)
			return
		}
		if user.Role < common.PERMISSION {
			ctx.JSON(errcode.NoPermission)
			return
		}
		ctx.Values().Set(common.ADMIN, true)
		ctx.Next()
	}
}
