package controller

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/service"
)

type postController struct {
	controller
}

func (c *postController) InsertPost(ctx iris.Context) {
	param := &vo.PostVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkEmpty(param.Content); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkUInt(param.CommentId); err != nil {
		ctx.JSON(err)
		return
	}
	user := c.GetUser(ctx)
	param.UserId = user.Id
	svc := service.NewPostService()
	result := svc.Insert(param)
	if result.IsSuccess() {
		c.Record(ctx, param.CommentId, common.KIND_CMNT, common.ACTION_CMNT)
	}
	ctx.JSON(result)
}

func (c *postController) UpdatePost(ctx iris.Context) {
	param := &vo.PostVo{}
	if err := c.getParams(ctx, param, false); err != nil {
		ctx.JSON(err)
		return
	}
	if err := c.checkUInt(param.Id); err != nil {
		ctx.JSON(err)
		return
	}
	svc := service.NewPostService()
	ctx.JSON(svc.Update(param))
}
