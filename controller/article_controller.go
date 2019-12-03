package controller

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/service"
)

type articleController struct {
	controller
}

func (c *articleController) QueryArticles(ctx iris.Context) {
	article := &vo.ArticleParam{}
	if err := c.getParams(ctx, article, true); err != nil {
		ctx.JSON(err)
		return
	}
	if c.IsAdminRouter(ctx) {
		article.Kind = 0
	} else {
		article.Kind = common.ARTICLE_NORMAL
		article.IsDelete = 1
	}
	svc := service.NewArticleService()
	ctx.JSON(svc.QueryArticles(article))
}

func (c *articleController) GetArticleById(ctx iris.Context) {
	id := ctx.URLParamIntDefault("id", 0)
	if id <= 0 {
		ctx.JSON(errcode.IllegalParam)
		return
	}
	if !c.IsAdminRouter(ctx) {
		go c.Record(ctx, id, common.KIND_ARTICLE, common.ACTION_VISIT)
	}
	svc := service.NewArticleService()
	ctx.JSON(svc.GetById(id, false))
}

func (c *articleController) InsertArticle(ctx iris.Context) {
	article := &vo.ArticleVo{}
	if err := c.getParams(ctx, article, true); err != nil {
		ctx.JSON(err)
		return
	}
	svc := service.NewArticleService()
	ctx.JSON(svc.Insert(article))
}

func (c *articleController) UpdateArticle(ctx iris.Context) {
	article := &vo.ArticleVo{}
	if err := c.getParams(ctx, article, false); err != nil {
		ctx.JSON(err)
		return
	}
	svc := service.NewArticleService()
	ctx.JSON(svc.Update(article))
}

func (c *articleController) DeleteArticle(ctx iris.Context) {
	article := &vo.ArticleDeleteParam{}
	if err := c.getParams(ctx, article, false); err != nil {
		ctx.JSON(err)
		return
	}
	svc := service.NewArticleService()
	ctx.JSON(svc.Delete(article))
}
