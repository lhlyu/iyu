package controller

import (
    "github.com/kataras/iris"
    "github.com/lhlyu/iyu/errcode"
    "github.com/lhlyu/iyu/service"
    "github.com/lhlyu/iyu/service/bo"
)

type articleController struct {
	controller
}

func (c *articleController) GetArticles(ctx iris.Context) {
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

type XX struct {
    A  string  `json:"a" form:"a"`
    B  int     `json:"b" form:"b"`
}

func (c *articleController) X(ctx iris.Context) {
    xx := &XX{}
    if err := c.getParams(ctx,xx,false);err != nil{
        ctx.JSON(err)
    }
    ctx.JSON(xx)
}
