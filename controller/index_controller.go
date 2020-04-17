package controller

import (
	"github.com/lhlyu/iyu/service"
	"github.com/lhlyu/yutil/v2"
)

type IndexController struct {
}

type HelloParam struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	M    []int  `json:"m"`
}

// http://localhost:8080/index?name=tom&age=12&m=1&m=2&m=3
func (c *IndexController) Hello(ctx *Context) {
	param := &HelloParam{}
	if !ctx.GetParams(param, false) {
		ctx.Info("12313")
		ctx.JSON("12312321")
		return
	}
	ctx.Info(yutil.Json.Marshal(param))
	svc := service.NewIndexService(ctx)
	ctx.JSON(svc.Hello(param.Name, param.Age))
}
