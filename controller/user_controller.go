package controller

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/service"
	"github.com/lhlyu/iyu/service/bo"
)

type userController struct {
	controller
}

func (c *userController) Login(ctx iris.Context) {
	id, _ := ctx.URLParamInt("id")
	if err := c.checkUInt(id); err != nil {
		ctx.JSON(err)
		return
	}
	user := service.NewUserService().GetById(id).Data.(*bo.UserData)
	xuser := &common.XUser{
		Id:   user.Id,
		Role: user.Role,
	}
	token := c.getToken(xuser)
	ctx.JSON(errcode.Success.WithData(token))
}

func (c *userController) UpdateUser(ctx iris.Context) {
	param := &vo.UserEditParam{}
	if err := c.getParams(ctx, param, false); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		ctx.JSON(err)
		return
	}
	ctx.JSON(service.NewUserService().Update(param))
}

func (c *userController) InsertUser(ctx iris.Context) {
	param := &vo.UserEditParam{}
	if err := c.getParams(ctx, param, false); err != nil {
		ctx.JSON(err)
		return
	}
	ctx.JSON(service.NewUserService().Insert(param))
}

func (c *userController) QueryUser(ctx iris.Context) {
	param := &vo.UserParam{}
	if err := c.getParams(ctx, param, true); err != nil {
		ctx.JSON(err)
		return
	}
	ctx.JSON(service.NewUserService().Query(param))
}
