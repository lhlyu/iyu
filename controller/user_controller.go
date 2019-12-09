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

/**
OAuth 协议的认证和授权的过程如下：

用户打开我的博客后，我想要通过GitHub获取改用户的基本信息
在转跳到GitHub的授权页面后，用户同意我获取他的基本信息
博客获得GitHub提供的授权码，使用该授权码向GitHub申请一个令牌
GitHub对博客提供的授权码进行验证，验证无误后，发放一个令牌给博客端
博客端使用令牌，向GitHub获取用户信息
GitHub 确认令牌无误，返回给我基本的用户信息
*/
func (c *userController) Login(ctx iris.Context) {
	id, _ := ctx.URLParamInt("id")
	if err := c.checkUInt(id); err != nil {
		c.Response(ctx, err)
		return
	}
	result := service.NewUserService(c.GetGUID(ctx)).Query(false, id)
	if !result.IsSuccess() {
		c.Response(ctx, result)
		return
	}
	users := result.Data.([]*bo.User)
	xuser := &common.XUser{
		Id:   users[0].Id,
		Role: users[0].Role,
	}
	token := c.getToken(xuser)
	c.Response(ctx, errcode.Success.WithData(token))
}

// todo
func (c *userController) GetAuthor(ctx iris.Context) {
	c.Response(ctx, service.NewUserService(c.GetGUID(ctx)).Get(1))
}

func (c *userController) UpdateUser(ctx iris.Context) {
	param := &vo.UserEditParam{}
	if err := c.getParams(ctx, param, false); err != nil {
		c.Response(ctx, err)
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		c.Response(ctx, err)
		return
	}
	c.Response(ctx, service.NewUserService(c.GetGUID(ctx)).Edit(param))
}

func (c *userController) InsertUser(ctx iris.Context) {
	param := &vo.UserEditParam{}
	if err := c.getParams(ctx, param, false); err != nil {
		c.Response(ctx, err)
		return
	}
	c.Response(ctx, service.NewUserService(c.GetGUID(ctx)).Edit(param))
}

func (c *userController) QueryUser(ctx iris.Context) {
	param := &vo.UserParam{}
	if err := c.getParams(ctx, param, true); err != nil {
		c.Response(ctx, err)
		return
	}
	c.Response(ctx, service.NewUserService(c.GetGUID(ctx)).QueryPage(param))
}
