package middleware

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/lhlyu/iyu/common"
	"time"
)

func Log() context.Handler {
	return func(ctx iris.Context) {
		start := time.Now()
		ctx.Next()
		reqInfo := fmt.Sprintf("%s,cost = %fs", ctx.String(), time.Now().Sub(start).Seconds())
		if common.Ylog != nil {
			common.Ylog.Debug(reqInfo)
		}
	}
}
