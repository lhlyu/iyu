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

func (c *tagController) GetTagAll(ctx iris.Context) {
	svc := service.NewTagService(c.GetGUID(ctx))
	c.Response(ctx, svc.Query(false))
}

func (c *tagController) UpdateTag(ctx iris.Context) {
	param := &vo.TagVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkUInt(param.IsDelete); err != nil {
		param.IsDelete = common.UNDELETED
	}
	if err := c.checkEmpty(param.Name); err != nil {
		c.Response(ctx, err)
		return
	}
	svc := service.NewTagService(c.GetGUID(ctx))
	c.Response(ctx, svc.Edit(param))
}

func (c *tagController) InsertTag(ctx iris.Context) {
	param := &vo.TagVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkEmpty(param.Name); err != nil {
		c.Response(ctx, err)
		return
	}
	svc := service.NewTagService(c.GetGUID(ctx))
	c.Response(ctx, svc.Edit(param))
}

func (c *tagController) DeleteTag(ctx iris.Context) {
	param := &vo.TagVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		c.Response(ctx, err)
		return
	}
	param.IsDelete = common.DELETED
	svc := service.NewTagService(c.GetGUID(ctx))
	c.Response(ctx, svc.Edit(param))
}
