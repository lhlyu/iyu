package model

import "time"

type YuArticle struct {
	Id             int       // 文章ID
	UserId         int       // 用户ID
	Like           int       // 赞
	Unlike         int       // 踩
	View           int       // 浏览量
	CommentsNumber int       // 评论数量
	Bg             string    // 头背景
	Title          string    // 标题
	Content        string    // 内容
	IsTop          int       // 是否置顶:0-不置顶;1-置顶
	NailId         int       // 钉子ID
	Kind           int       // 文章类型:0-普通文章;1-特殊文章
	IsDelete       int       // 是否已删除:0-未删除;1-已删除
	CreatedAt      time.Time // 创建时间
	UpdatedAt      time.Time // 更新时间
}
