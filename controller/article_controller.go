package controller

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/controller/dto"
	"github.com/lhlyu/iyu/service/article_service"
)

type articleController struct {
	controller
}

// 获取首页文章列表
func (c *articleController) GetHomeArticlePage(ctx iris.Context) {
	param := &dto.ArticleDto{}
	if !c.getParams(ctx, param, true) {
		return
	}
	param.IsDelete = 1
	param.Kind = 1
	svc := article_service.NewService(c.GetTraceId(ctx))
	ctx.JSON(svc.QueryHomeArticlePage(param))
}

// 用户获取单篇文章
func (c *articleController) GetArticleByCode(ctx iris.Context) {
	param := &dto.ArticleDto{}
	if !c.getParams(ctx, param, false) {
		return
	}
	if err := c.checkEmpty(param.Code); err != nil {
		ctx.JSON(err)
		return
	}
	svc := article_service.NewService(c.GetTraceId(ctx))
	ctx.JSON(svc.GetArticleByCode(param.Code))
}

// 获取about文章
func (c *articleController) GetAbout(ctx iris.Context) {
	svc := article_service.NewService(c.GetTraceId(ctx))
	ctx.JSON(svc.GetAbout())
}

// 获取时间线
func (c *articleController) GetTimeline(ctx iris.Context) {
	svc := article_service.NewService(c.GetTraceId(ctx))
	ctx.JSON(svc.GetTimeline())
}

// 管理页获取文章列表
func (c *articleController) GetAdminArticlePage(ctx iris.Context) {
	param := &dto.ArticleDto{}
	if !c.getParams(ctx, param, true) {
		return
	}
	svc := article_service.NewService(c.GetTraceId(ctx))
	ctx.JSON(svc.QueryAdminArticlePage(param))
}

// 管理获取单篇文章
func (c *articleController) GetArticleById(ctx iris.Context) {
	param := &dto.ArticleDto{}
	if !c.getParams(ctx, param, false) {
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		ctx.JSON(err)
		return
	}
	svc := article_service.NewService(c.GetTraceId(ctx))
	ctx.JSON(svc.GetArticleById(param.Id))
}

// 添加文章
func (*articleController) AddArticle(ctx iris.Context) {

}

// 修改文章
func (*articleController) UpdateArticle(ctx iris.Context) {

}
