package controller

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/service"
)

type tagController struct {
	controller
}

func (*tagController) GetTagAll(ctx iris.Context) {
	svc := service.NewTagService()
	ctx.JSON(svc.GetAll(false))
}

func (c *tagController) UpdateTag(ctx iris.Context) {
	param := &vo.TagVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkUInt(param.IsDelete); err != nil {
		param.IsDelete = common.UNDELETED
	}
	if err := c.checkEmpty(param.Name); err != nil {
		ctx.JSON(err)
		return
	}
	svc := service.NewTagService()
	ctx.JSON(svc.Update(param))
}

func (c *tagController) InsertTag(ctx iris.Context) {
	param := &vo.TagVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkEmpty(param.Name); err != nil {
		ctx.JSON(err)
		return
	}
	svc := service.NewTagService()
	ctx.JSON(svc.Insert(param))
}

func (c *tagController) DeleteTag(ctx iris.Context) {
	param := &vo.TagVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		ctx.JSON(err)
		return
	}
	param.IsDelete = common.DELETED
	svc := service.NewTagService()
	ctx.JSON(svc.Delete(param))
}
