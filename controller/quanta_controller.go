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
		c.Response(ctx, err)
		return
	}
	svc := service.NewQuantaService(c.GetGUID(ctx))
	c.Response(ctx, svc.QueryPage(page))
}

func (c *quantaController) UpdateQuanta(ctx iris.Context) {
	param := &vo.QuantaVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkEmpty(param.Key); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkEmpty(param.Value); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkUInt(param.IsEnable); err != nil {
		param.IsEnable = 1
	}
	svc := service.NewQuantaService(c.GetGUID(ctx))
	c.Response(ctx, svc.Edit(param))
}

func (c *quantaController) InsertQuanta(ctx iris.Context) {
	param := &vo.QuantaVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkEmpty(param.Key); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkEmpty(param.Value); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkUInt(param.IsEnable); err != nil {
		param.IsEnable = 1
	}
	svc := service.NewQuantaService(c.GetGUID(ctx))
	c.Response(ctx, svc.Edit(param))
}

func (c *quantaController) DeleteQuanta(ctx iris.Context) {
	param := &vo.QuantaVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		c.Response(ctx, err)
		return
	}
	param.IsEnable = common.DELETED
	svc := service.NewQuantaService(c.GetGUID(ctx))
	c.Response(ctx, svc.Edit(param))
}
