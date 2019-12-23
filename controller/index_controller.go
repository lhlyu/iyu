package controller

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/errcode"
)

type indexController struct {
	controller
}

// 获取网站信息(作者,背景图,统计...)
func (*indexController) GetWebSiteOption(ctx iris.Context) {

}

// 修改网站设置
func (*indexController) UpdateWebSiteOption(ctx iris.Context) {

}

// 网站操作记录
func (c *indexController) GetWebSiteLog(ctx iris.Context) {
}

// 处理请求错误
func (c *indexController) NoFoundHandler(ctx iris.Context) {
	ctx.JSON(errcode.NofoundError)
}

// 用户动作
func (c *indexController) UserAction(ctx iris.Context) {
}
