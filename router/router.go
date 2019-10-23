package router

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/lhlyu/iyu/repository"
)

func SetRouter(app *iris.Application) {
	// only a test
	app.Get("/", func(ctx context.Context) {
		d := repository.NewDao()
		data, e := d.QueryNail()
		if e != nil {
			fmt.Println(e)
		}
		ctx.JSON(data)
	})
}
