package controller

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
)

type UserController struct {
	controller
}

func (c *UserController) GetToken(ctx iris.Context) {
	m := map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	ctx.JSON(c.getToken(m))
}

func (c *UserController) GetToken2(ctx iris.Context) {
	user, ok := ctx.Values().Get("jwt").(*jwt.Token)
	if ok {
		ctx.WriteString("error")
		return
	}

	foobar := user.Claims.(jwt.MapClaims)
	for key, value := range foobar {
		ctx.Writef("%s = %v\n", key, value)
	}

}
