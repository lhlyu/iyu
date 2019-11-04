package router

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository"
)

func SetRouter(app *iris.Application) {

	app.AllowMethods(iris.MethodOptions)
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(errcode.Error.AddMsg("not found resources"))
	})

	// only a test
	app.Get("/", func(ctx iris.Context) {
		golog.Debug(ctx.String())
		d := repository.NewDao()
		data, e := d.QueryNail()
		if e != nil {
			fmt.Println(e)
		}
		ctx.JSON(data)
	})
}
