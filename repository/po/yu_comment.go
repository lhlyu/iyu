package model

import "time"

type YuComment struct {
	Id        int       `json:"id"`        // 主键ID
	ArticleId int       `json:"articleId"` // 文章ID
	UserId    int       `json:"userId"`    // 用户ID
	Floor     string    `json:"floor"`     // 楼层
	Content   string    `json:"content"`   // 评论内容
	Like      int       `json:"like"`      // 赞
	Unlike    int       `json:"unlike"`    // 踩
	IsCheck   int       `json:"isCheck"`   // 评论是否已审核:0-未审核;1-已审核
	IsDelete  int       `json:"isDelete"`  // 评论是否已被删除:0-未删除;1-已删除
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 修改时间
}
