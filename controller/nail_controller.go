package controller

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/service"
)

type nailController struct {
	controller
}

func (*nailController) GetNailAll(ctx iris.Context) {
	svc := service.NewNailService()
	ctx.JSON(svc.GetAll(false))
}

func (c *nailController) UpdateNail(ctx iris.Context) {
	param := &vo.NailVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkEmpty(param.Name); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkUInt(param.Status); err != nil {
		param.Status = 1
	}
	if err := c.checkEmpty(param.Color); err != nil {
		param.Color = "#000000"
	}
	svc := service.NewNailService()
	ctx.JSON(svc.Update(param.Id, param.Status, param.Name, param.Color))
}

func (c *nailController) InsertNail(ctx iris.Context) {
	param := &vo.NailVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkEmpty(param.Name); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkEmpty(param.Color); err != nil {
		param.Color = "#000000"
	}
	svc := service.NewNailService()
	ctx.JSON(svc.Insert(param.Name, param.Color))
}

func (c *nailController) DeleteNail(ctx iris.Context) {
	param := &vo.NailVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		ctx.JSON(err)
		return
	}
	svc := service.NewNailService()
	ctx.JSON(svc.Delete(param.Id, param.Real))
}
