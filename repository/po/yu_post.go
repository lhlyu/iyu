package model

import "time"

type YuPost struct {
	Id        int       `json:"id"`        // 主键ID
	CommentId int       `json:"commentId"` // 评论ID
	UserId    int       `json:"userId"`    // 用户ID
	Floor     string    `json:"floor"`     // 楼层
	AtId      int       `json:"atId"`      // 艾特回复的ID
	AtFloor   string    `json:"atFloor"`   // 艾特回复的楼层
	Content   string    `json:"content"`   // 评论内容
	Like      int       `json:"like"`      // 赞
	Unlike    int       `json:"unlike"`    // 踩
	IsCheck   int       `json:"isCheck"`   // 评论是否已审核:0-未审核;1-已审核
	IsDelete  int       `json:"isDelete"`  // 评论是否已被删除:0-未删除;1-已删除
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 修改时间
}
