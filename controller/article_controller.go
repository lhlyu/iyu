package controller

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/service"
	"github.com/lhlyu/iyu/service/bo"
)

type articleController struct {
	controller
}

func (c *articleController) GetArticles(ctx iris.Context) {
	common.Ylog.Debug("GetArticles")
	param := &bo.ArticleParam{}
	if err := c.getParams(ctx, param, true); err != nil {
		ctx.JSON(err)
		return
	}
	svc := service.NewArticleService()
	result := svc.GetArticles(param)
	ctx.JSON(result)
}

func (*articleController) GetArticleById(ctx iris.Context) {
	id, e := ctx.Params().GetInt("id")
	if e != nil {
		ctx.JSON(errcode.IllegalParam)
		return
	}
	ctx.JSON(id)
}
