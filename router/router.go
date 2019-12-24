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
	user := api.Party("/user", middleware.PermissionUser())
	{
		user.Post("/cmnt", ctr.AddCmnt)
		user.Post("/reply", ctr.AddReply)
	}

	// 针对管理员
	admin := api.Party("/admin", middleware.PermissionAdmin())
	{
		admin.Get("/articles", ctr.GetAdminArticlePage)
		admin.Get("/article", ctr.GetArticleById)
		admin.Post("/article", ctr.AddArticle)
		admin.Put("/article", ctr.UpdateArticle)

		admin.Get("/tags", ctr.GetTagPage)
		admin.Post("/tag", ctr.AddTag)
		admin.Put("/tag", ctr.UpdateTag)
		admin.Delete("/tag", ctr.BatchDeleteTag)

		admin.Get("/categorys", ctr.GetCategoryPage)
		admin.Post("/category", ctr.UpdateCategory)
		admin.Put("/category", ctr.AddCategory)
		admin.Delete("/category", ctr.BatchDeleteCategory)

		admin.Get("/quantas", ctr.GetQuantaPage)
		admin.Put("/quanta", ctr.UpdateQuanta)

		admin.Get("/users", ctr.GetUserPage)
		admin.Get("/user", ctr.GetUserById)
		admin.Post("/user", ctr.UpdateUser)

		admin.Get("/cmnt", ctr.GetCmntAndReplyPage)
		admin.Post("/cmnt", ctr.UpdateCmnt)

		admin.Post("/reply", ctr.UpdateReply)

		admin.Get("/logs", ctr.GetWebSiteLog)

		admin.Put("/setting", ctr.UpdateWebSiteOption)
	}
}
