package router

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/controller"
)

func SetRouter(app *iris.Application) {
	app.AllowMethods(iris.MethodOptions)

	ctr := &controller.Controller{}

	app.OnErrorCode(iris.StatusNotFound, ctr.NoFoundHandler)

	api := app.Party("/api")
	{
		api.Get("/", ctr.Login)
	}
	//{
	//	api.Use(middleware.PermissionUser())
	//}
	//api.Party("/admin", middleware.PermissionAdmin())
	//{
	//
	//
	//}
}
