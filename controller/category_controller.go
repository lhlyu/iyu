package controller

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/service"
)

type categoryController struct {
	controller
}

func (c *categoryController) GetCategoryAll(ctx iris.Context) {
	svc := service.NewCategoryService(c.GetGUID(ctx))
	c.Response(ctx, svc.Query(false))
}

func (c *categoryController) UpdateCategory(ctx iris.Context) {
	param := &vo.CategoryVo{}
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
	svc := service.NewCategoryService(c.GetGUID(ctx))
	c.Response(ctx, svc.Edit(param))
}

func (c *categoryController) InsertCategory(ctx iris.Context) {
	param := &vo.CategoryVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkEmpty(param.Name); err != nil {
		c.Response(ctx, err)
		return
	}
	svc := service.NewCategoryService(c.GetGUID(ctx))
	c.Response(ctx, svc.Edit(param))
}

func (c *categoryController) DeleteCategory(ctx iris.Context) {
	param := &vo.CategoryVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		c.Response(ctx, err)
		return
	}
	param.IsDelete = common.DELETED
	svc := service.NewCategoryService(c.GetGUID(ctx))
	c.Response(ctx, svc.Edit(param))
}
