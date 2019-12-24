package middleware

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/util"
)

func Log() context.Handler {
	return func(ctx iris.Context) {
		// 加入唯一ID
		traceId := util.GetGID()
		ctx.Values().Set(common.X_TRACE, traceId)
		common.Ylog.Log(2, "debug", traceId, "middleware", ctx.String())
		ctx.Next()
	}
}
