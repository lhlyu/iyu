package middleware

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/util"
	"time"
)

func Log() context.Handler {
	return func(ctx iris.Context) {
		start := time.Now()
		ctx.Next()
		requestInfo := fmt.Sprintf("%s ▶ %s:%s", util.RemoteIp(ctx.Request()), ctx.Method(), ctx.Request().RequestURI)
		reqInfo := fmt.Sprintf("%s,cost = %fs", requestInfo, time.Now().Sub(start).Seconds())
		if common.Ylog != nil {
			common.Ylog.Debug(reqInfo)
		}
	}
}
