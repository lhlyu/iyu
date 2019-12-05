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

func (*nailController) GetNailAll(ctx iris.Context) {
	svc := service.NewNailService()
	ctx.JSON(svc.Query(false))
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
	if err := c.checkUInt(param.IsDelete); err != nil {
		param.IsDelete = common.UNDELETED
	}
	if err := c.checkEmpty(param.Color); err != nil {
		param.Color = common.COLOR
	}
	svc := service.NewNailService()
	ctx.JSON(svc.Edit(param))
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
		param.Color = common.COLOR
	}
	svc := service.NewNailService()
	ctx.JSON(svc.Edit(param))
}

// id real
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
	param.IsDelete = common.DELETED
	svc := service.NewNailService()
	ctx.JSON(svc.Edit(param))
}
