package controller

import (
	"github.com/lhlyu/iyu/service"
)

type CategoryController struct {
}

func (c CategoryController) Query(ctx *Context) {
	svc := service.NewCategoryService(ctx.GetTracker())
	svc.GetOne()
	ctx.Info(">>>>>>>>>>>>")
	ctx.JSON("")
}
