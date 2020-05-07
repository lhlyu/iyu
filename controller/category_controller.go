package controller

import (
	"github.com/lhlyu/iyu/service"
)

type CategoryController struct {
}

func (c CategoryController) Query(ctx *Context) {
	svc := service.NewCategoryService(ctx.GetTracker())
	svc.GetAll()
	ctx.Info(">>>>>>>>>>>>")
	ctx.JSON("")
}
