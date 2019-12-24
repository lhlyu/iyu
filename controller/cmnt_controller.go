package controller

import "github.com/kataras/iris"

type cmntController struct {
	controller
}

// 用户获取评论列表
func (*cmntController) GetCmntPage(ctx iris.Context) {

}

// 发表评论
func (*cmntController) AddCmnt(ctx iris.Context) {

}

// 后台获取评论回复列表
func (*cmntController) GetCmntAndReplyPage(ctx iris.Context) {

}

// 修改评论
func (*cmntController) UpdateCmnt(ctx iris.Context) {

}
