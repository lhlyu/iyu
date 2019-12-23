package router

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/controller"
	"github.com/lhlyu/iyu/middleware"
)

func SetRouter(app *iris.Application) {
	app.AllowMethods(iris.MethodOptions)

	ctr := &controller.Controller{}

	app.OnErrorCode(iris.StatusNotFound, ctr.NoFoundHandler)

	api := app.Party("/api")

	// 针对游客
	{
		api.Get("/login", ctr.Login)
		api.Get("/articles", ctr.GetHomeArticlePage)
		api.Get("/article", ctr.GetArticleByCode)
		api.Get("/about", ctr.GetAbout)
		api.Get("/timeline", ctr.GetTimeline)
		api.Get("/action", ctr.UserAction)
		api.Get("/website", ctr.GetWebSiteOption)
		api.Get("/cmnts", ctr.GetCmntPage)
		api.Get("/reply", ctr.GetReplyPage)
	}

	// 针对已登录的用户
	api.Use(middleware.PermissionUser())
	{
		api.Post("/cmnt", ctr.AddCmnt)
		api.Post("/reply", ctr.AddReply)
	}

	// 针对管理员
	api.Use(middleware.PermissionAdmin())
	api.Party("/admin")
	{
		api.Get("/articles", ctr.GetAdminArticlePage)
		api.Get("/article", ctr.GetArticleById)
		api.Post("/article", ctr.AddArticle)
		api.Put("/article", ctr.UpdateArticle)

		api.Get("/tags", ctr.GetTagPage)
		api.Post("/tag", ctr.AddTag)
		api.Put("/tag", ctr.UpdateTag)
		api.Delete("/tag", ctr.BatchDeleteTag)

		api.Get("/categorys", ctr.GetCategoryPage)
		api.Post("/category", ctr.UpdateCategory)
		api.Put("/category", ctr.AddCategory)
		api.Delete("/category", ctr.BatchDeleteCategory)

		api.Get("/quantas", ctr.GetQuantaPage)
		api.Put("/quanta", ctr.UpdateQuanta)

		api.Get("/users", ctr.GetUserPage)
		api.Get("/user", ctr.GetUserById)
		api.Post("/user", ctr.UpdateUser)

		api.Get("/cmnt", ctr.GetCmntAndReplyPage)
		api.Post("/cmnt", ctr.UpdateCmnt)

		api.Post("/reply", ctr.UpdateReply)

		api.Get("/logs", ctr.GetWebSiteLog)

		api.Put("/setting", ctr.UpdateWebSiteOption)
	}
}
