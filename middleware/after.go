package middleware

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

// 可以做统一响应处理 比如打印日志记录响应结果...
func After() context.Handler {
	return func(ctx iris.Context) {
		//body := ctx.Recorder().Body()  // 获取响应返回的内容
		//ctx.Recorder().ResetBody()     // 将响应体body内容置空
		//traceId := ctx.Values().Get(common.X_TRACE).(string)
		//common.Ylog.Log(2,"info",traceId,"middleware",string(body))
		//ctx.Write(body)                // 重写写入body
	}
}
