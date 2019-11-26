package controller

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/service"
)

type categoryController struct {
	controller
}

func (*categoryController) GetCategoryAll(ctx iris.Context) {
	svc := service.NewCategoryService()
	ctx.JSON(svc.GetAll(false))
}

func (c *categoryController) UpdateCategory(ctx iris.Context) {
	param := &vo.CategoryVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkUInt(param.Status); err != nil {
		param.Status = 1
	}
	if err := c.checkEmpty(param.Name); err != nil {
		ctx.JSON(err)
		return
	}
	svc := service.NewCategoryService()
	ctx.JSON(svc.Update(param.Id, param.Status, param.Name))
}

func (c *categoryController) InsertCategory(ctx iris.Context) {
	param := &vo.CategoryVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkEmpty(param.Name); err != nil {
		ctx.JSON(err)
		return
	}
	svc := service.NewCategoryService()
	ctx.JSON(svc.Insert(param.Name))
}

func (c *categoryController) DeleteCategory(ctx iris.Context) {
	param := &vo.CategoryVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		ctx.JSON(err)
		return
	}
	svc := service.NewCategoryService()
	ctx.JSON(svc.Delete(param.Id, param.Real))
}
