package middleware

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/lhlyu/iyu/common"
)

func Jwt() context.Handler {
	return func(ctx iris.Context) {
		user := &common.XUser{}
		user.Ip = ctx.RemoteAddr()
		if err := jwt.New(jwt.Config{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte(common.Cfg.GetString("jwt.secret")), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
		}).CheckJWT(ctx); err == nil {
			tokens, _ := ctx.Values().Get("jwt").(*jwt.Token)
			tokenMap := tokens.Claims.(jwt.MapClaims)
			user.Id = tokenMap[common.X_ID].(int)
			user.Role = tokenMap[common.X_ROLE].(int)
		}
		ctx.Values().Set(common.X_USER, user)
		ctx.Next()
	}
}
