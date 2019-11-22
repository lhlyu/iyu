package controller

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/service"
)

type tagController struct {
	controller
}

func (*tagController) GetTagAll(ctx iris.Context) {
	svc := service.NewTagService()
	result := svc.GetTagAll()
	ctx.JSON(result)
}
