package controller

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/dto"
	"github.com/lhlyu/iyu/service/quanta_service"
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
	svc := quanta_service.NewService(c.GetTraceId(ctx))
	svc.LoadQuanta(0)
	param := &dto.QuantaDto{
		Page: common.NewPageAll(),
	}
	ctx.JSON(svc.QueryQuantaPage(param))
}
