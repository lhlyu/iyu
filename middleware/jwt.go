package middleware

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/context"
	"github.com/lhlyu/iyu/common"
)

func Jwt() context.Handler {
	return jwt.New(jwt.Config{
		Extractor: jwt.FromParameter("token"),
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(common.Cfg.GetString("jwt.secret")), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	}).Serve
}
