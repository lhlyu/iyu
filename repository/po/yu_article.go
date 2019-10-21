package model

import "time"

type YuArticle struct {
	Id             int       `json:"id"`             // 文章ID
	UserId         int       `json:"userId"`         // 用户ID
	Like           int       `json:"like"`           // 赞
	Unlike         int       `json:"unlike"`         // 踩
	View           int       `json:"view"`           // 浏览量
	CommentsNumber int       `json:"commentsNumber"` // 评论数量
	Bg             string    `json:"bg"`             // 头背景
	Title          string    `json:"title"`          // 标题
	Content        string    `json:"content"`        // 内容
	IsTop          int       `json:"isTop"`          // 是否置顶:0-不置顶;1-置顶
	NailId         int       `json:"nailId"`         // 钉子ID
	Kind           int       `json:"kind"`           // 文章类型:0-普通文章;1-特殊文章
	IsDelete       int       `json:"isDelete"`       // 是否已删除:0-未删除;1-已删除
	CreatedAt      time.Time `json:"createdAt"`      // 创建时间
	UpdatedAt      time.Time `json:"updatedAt"`      // 更新时间
}
