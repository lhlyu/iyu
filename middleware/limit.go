package middleware

import (
	"github.com/didip/tollbooth"
	"github.com/iris-contrib/middleware/tollboothic"
	"github.com/kataras/iris/context"
	"github.com/lhlyu/iyu/common"
)

// 限制每秒请求数量
func Limiter() context.Handler {
	limiter := tollbooth.NewLimiter(common.Cfg.GetFloat64("server.limit"), nil)
	return tollboothic.LimitHandler(limiter)
}
