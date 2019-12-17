package middleware

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/yutil"
)

/**
Authorization:bearer xxxxxxxxxxx
*/
func Jwt() context.Handler {
	return func(ctx iris.Context) {
		traceId := ctx.Values().Get(common.X_TRACE).(string)
		user := &common.XUser{}
		user.Ip = yutil.ClientIp(ctx.Request())
		var err error
		if err = jwt.New(jwt.Config{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte(common.Cfg.GetString("jwt.secret")), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
			Expiration:    true,
		}).CheckJWT(ctx); err == nil {
			token, _ := jwt.FromAuthHeader(ctx)
			tokens, _ := ctx.Values().Get("jwt").(*jwt.Token)
			tokenMap := tokens.Claims.(jwt.MapClaims)
			user.Id = int(tokenMap[common.X_ID].(float64))
			user.Role = int(tokenMap[common.X_ROLE].(float64))
			cacheToken := cache.NewCache(traceId).GetJwt(user.Id)
			if cacheToken == "" || token != cacheToken {
				ctx.JSON(errcode.NoPermission)
				return
			}
		}
		ctx.Values().Set(common.X_USER, user)
		ctx.Next()
	}
}
