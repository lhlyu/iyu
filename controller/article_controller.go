package controller

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/controller/vo"
)

type articleController struct {
	controller
}

func (c *articleController) InsertArticle(ctx iris.Context) {
	article := &vo.ArticleVo{}
	if err := c.getParams(ctx, article, true); err != nil {
		ctx.JSON(err)
		return
	}
}
