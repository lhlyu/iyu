package middleware

import (
	"fmt"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/lhlyu/iyu/common"
)

/**
Authorization:bearer xxxxxxxxxxx
*/
func Jwt() context.Handler {
	return func(ctx iris.Context) {
		user := &common.XUser{}
		user.Ip = ctx.RemoteAddr()
		var err error
		if err = jwt.New(jwt.Config{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte(common.Cfg.GetString("jwt.secret")), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
		}).CheckJWT(ctx); err == nil {
			tokens, _ := ctx.Values().Get("jwt").(*jwt.Token)
			tokenMap := tokens.Claims.(jwt.MapClaims)
			user.Id = int(tokenMap[common.X_ID].(float64))
			user.Role = int(tokenMap[common.X_ROLE].(float64))
		}
		fmt.Println("jwt.err = ", err)
		ctx.Values().Set(common.X_USER, user)
		ctx.Next()
	}
}
