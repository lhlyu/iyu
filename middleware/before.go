package middleware

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func Before() context.Handler {
	return func(ctx iris.Context) {
		ctx.Record() // 开启响应body记录
		ctx.Next()
	}
}
