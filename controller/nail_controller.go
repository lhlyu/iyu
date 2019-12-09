package controller

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/service"
)

type nailController struct {
	controller
}

func (c *nailController) GetNailAll(ctx iris.Context) {
	svc := service.NewNailService(c.GetGUID(ctx))
	c.Response(ctx, svc.Query(false))
}

func (c *nailController) UpdateNail(ctx iris.Context) {
	param := &vo.NailVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkEmpty(param.Name); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkUInt(param.IsDelete); err != nil {
		param.IsDelete = common.UNDELETED
	}
	if err := c.checkEmpty(param.Color); err != nil {
		param.Color = common.COLOR
	}
	svc := service.NewNailService(c.GetGUID(ctx))
	c.Response(ctx, svc.Edit(param))
}

func (c *nailController) InsertNail(ctx iris.Context) {
	param := &vo.NailVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkEmpty(param.Name); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkEmpty(param.Color); err != nil {
		param.Color = common.COLOR
	}
	svc := service.NewNailService(c.GetGUID(ctx))
	c.Response(ctx, svc.Edit(param))
}

// id real
func (c *nailController) DeleteNail(ctx iris.Context) {
	param := &vo.NailVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		c.Response(ctx, err)
		return
	}
	param.IsDelete = common.DELETED
	svc := service.NewNailService(c.GetGUID(ctx))
	c.Response(ctx, svc.Edit(param))
}
