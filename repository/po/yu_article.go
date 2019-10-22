package model

import "time"

type YuArticle struct {
	Id             int       `db:"id"`              // 文章ID
	UserId         int       `db:"user_id"`         // 用户ID
	Like           int       `db:"like"`            // 赞
	Unlike         int       `db:"unlike"`          // 踩
	View           int       `db:"view"`            // 浏览量
	CommentsNumber int       `db:"comments_number"` // 评论数量
	Bg             string    `db:"bg"`              // 头背景
	Title          string    `db:"title"`           // 标题
	Content        string    `db:"content"`         // 内容
	IsTop          int       `db:"is_top"`          // 是否置顶:0-不置顶;1-置顶
	NailId         int       `db:"nail_id"`         // 钉子ID
	Kind           int       `db:"kind"`            // 文章类型:0-普通文章;1-特殊文章
	IsDelete       int       `db:"is_delete"`       // 是否已删除:0-未删除;1-已删除
	CreatedAt      time.Time `db:"created_at"`      // 创建时间
	UpdatedAt      time.Time `db:"updated_at"`      // 更新时间
}
