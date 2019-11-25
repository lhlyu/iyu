package middleware

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

// 记录器 浏览统计 用户动作记录 todo
func Record() context.Handler {
	return func(ctx iris.Context) {
		user, ok := ctx.Values().Get("jwt").(*jwt.Token)
		if ok {
			ctx.WriteString("error")
			return
		}

		foobar := user.Claims.(jwt.MapClaims)
		for key, value := range foobar {
			ctx.Writef("%s = %v\n", key, value)
		}
		ctx.Next()
	}
}
