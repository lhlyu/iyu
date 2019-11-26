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

		// tag
		api.Post("/tag", ctr.InsertTag)
		api.Delete("/tag", ctr.DeleteTag)
		api.Put("/tag", ctr.UpdateTag)
		api.Get("/tag", ctr.GetTagAll)
		// category
		api.Post("/category", ctr.InsertCategory)
		api.Delete("/category", ctr.DeleteCategory)
		api.Put("/category", ctr.UpdateCategory)
		api.Get("/category", ctr.GetCategoryAll)
		// nail
		api.Post("/nail", ctr.InsertNail)
		api.Delete("/nail", ctr.DeleteNail)
		api.Put("/nail", ctr.UpdateNail)
		api.Get("/nail", ctr.GetNailAll)
		// quanta
		api.Post("/quanta", ctr.InsertQuanta)
		api.Delete("/quanta", ctr.DeleteQuanta)
		api.Put("/quanta", ctr.UpdateQuanta)
		api.Get("/quanta", ctr.GetQuantaAll)

		api.Post("/x", ctr.X)
	}
}
