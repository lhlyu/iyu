package router

import (
	"github.com/kataras/iris/v12"
	"github.com/lhlyu/iyu/controller"
)

func SetRouter(app *iris.Application) {
	//app.AllowMethods(iris.MethodOptions)
	//
	//ctr := &controller.Controller{}
	//
	//
	//app.Party("/api")

	categoryCtr := controller.CategoryController{}

	api := app.Party("/api")

	api.Get("/category", controller.H(categoryCtr.Query))

}
