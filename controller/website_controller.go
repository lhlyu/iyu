package controller

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/errcode"
)

type websiteController struct {
	controller
}

func (c *websiteController) NoFoundHandler(ctx iris.Context) {
	ctx.Values().Set("resp", errcode.NofoundError)
}
