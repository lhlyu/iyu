package controller

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/service"
)

type quantaController struct {
	controller
}

func (c *quantaController) GetQuantaAll(ctx iris.Context) {
	page := &common.Page{}
	if err := c.getParams(ctx, page, true); err != nil {
		ctx.JSON(err)
		return
	}
	svc := service.NewQuantaService()
	ctx.JSON(svc.QueryPage(page))
}

func (c *quantaController) UpdateQuanta(ctx iris.Context) {
	param := &vo.QuantaVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkEmpty(param.Key); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkEmpty(param.Value); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkUInt(param.IsEnable); err != nil {
		param.IsEnable = 1
	}
	svc := service.NewQuantaService()
	ctx.JSON(svc.Edit(param))
}

func (c *quantaController) InsertQuanta(ctx iris.Context) {
	param := &vo.QuantaVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkEmpty(param.Key); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkEmpty(param.Value); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkUInt(param.IsEnable); err != nil {
		param.IsEnable = 1
	}
	svc := service.NewQuantaService()
	ctx.JSON(svc.Edit(param))
}

func (c *quantaController) DeleteQuanta(ctx iris.Context) {
	param := &vo.QuantaVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		ctx.JSON(err)
		return
	}
	param.IsEnable = common.DELETED
	svc := service.NewQuantaService()
	ctx.JSON(svc.Edit(param))
}
