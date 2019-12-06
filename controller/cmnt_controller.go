package controller

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/service"
)

type cmntController struct {
	controller
}

func (c *cmntController) InsertCmnt(ctx iris.Context) {
	param := &vo.CmntVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkEmpty(param.Content); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkUInt(param.ArticleId); err != nil {
		ctx.JSON(err)
		return
	}
	user := c.GetUser(ctx)
	param.UserId = user.Id
	svc := service.NewCmntService()
	result := svc.Insert(param)
	if result.IsSuccess() {
		c.Record(ctx, param.ArticleId, common.KIND_ARTICLE, common.ACTION_CMNT)
	}
	ctx.JSON(result)
}

func (c *cmntController) UpdateCmnt(ctx iris.Context) {
	param := &vo.CmntVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		ctx.JSON(err)
		return
	}
	svc := service.NewCmntService()
	ctx.JSON(svc.Update(param))
}
