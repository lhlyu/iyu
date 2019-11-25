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

	ctr := &controller.Controller{}

	api := app.Party("/api")
	{

		api.Get("/", ctr.GetToken)
		api.Get("/p", middleware.Jwt(), ctr.GetToken2)

		api.Get("/articles", ctr.GetArticles)
		api.Get("/articles/{id:int}", ctr.GetArticleById)

		api.Post("/tag", ctr.InsertTag)
		api.Delete("/tag", ctr.DeleteTag)
		api.Put("/tag", ctr.UpdateTag)
		api.Get("/tag", ctr.GetTagAll)

		api.Post("/x", ctr.X)
	}
}
