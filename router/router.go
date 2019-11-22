package router

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/controller"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/middleware"
)

func SetRouter(app *iris.Application) {

	app.AllowMethods(iris.MethodOptions)
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(errcode.Error.AddMsg("not found resources"))
	})

	api := app.Party("/api")
	{
		ctr := &controller.Controller{}

		api.Get("/", ctr.GetToken)
		api.Get("/p", middleware.Jwt(), ctr.GetToken2)

		api.Get("/articles", ctr.GetArticles)
		api.Get("/articles/{id:int}", ctr.GetArticleById)
		//api.Get("/author")
		//api.Get("/website")
		//api.Get("/categorys")
		api.Get("/tags", ctr.GetTagAll)

		api.Post("/x", ctr.X)
	}
}
