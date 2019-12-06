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
		ctx.JSON(errcode.NofoundError)
		return
	})

	ctr := &controller.Controller{}

	api := app.Party("/api")
	{
		api.Get("/login", ctr.Login)
		// article
		api.Get("/articles", ctr.QueryArticles)
		api.Get("/article", ctr.GetArticleById)
	}
	{
		api.Use(middleware.PermissionUser())
		api.Post("/cmnt", ctr.InsertCmnt)
		api.Post("/post", ctr.InsertPost)
	}
	admin := api.Party("/admin", middleware.PermissionAdmin())
	{
		api.Put("/cmnt", ctr.UpdateCmnt)
		api.Put("/post", ctr.UpdatePost)

		admin.Get("/articles", ctr.QueryArticles)
		admin.Get("/article", ctr.GetArticleById)
		admin.Post("/article", ctr.InsertArticle)
		admin.Put("/article", ctr.UpdateArticle)
		admin.Delete("/article", ctr.DeleteArticle)

		admin.Get("/quanta", ctr.GetQuantaAll)
		admin.Post("/quanta", ctr.InsertQuanta)
		admin.Put("/quanta", ctr.UpdateQuanta)
		admin.Delete("/quanta", ctr.DeleteQuanta)

		admin.Get("/nail", ctr.GetNailAll)
		admin.Post("/nail", ctr.InsertNail)
		admin.Put("/nail", ctr.UpdateNail)
		admin.Delete("/nail", ctr.DeleteNail)

		admin.Get("/nail", ctr.GetNailAll)
		admin.Post("/nail", ctr.InsertNail)
		admin.Put("/nail", ctr.UpdateNail)
		admin.Delete("/nail", ctr.DeleteNail)

		admin.Get("/category", ctr.GetCategoryAll)
		admin.Post("/category", ctr.InsertCategory)
		admin.Put("/category", ctr.UpdateCategory)
		admin.Delete("/category", ctr.DeleteCategory)

		admin.Get("/tag", ctr.GetTagAll)
		admin.Post("/tag", ctr.InsertTag)
		admin.Put("/tag", ctr.UpdateTag)
		admin.Delete("/tag", ctr.DeleteTag)

		admin.Get("/users", ctr.QueryUser)
		admin.Put("/user", ctr.UpdateUser)
	}
}
