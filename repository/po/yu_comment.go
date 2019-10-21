package model

import "time"

type YuComment struct {
	Id        int       // 主键ID
	ArticleId int       // 文章ID
	UserId    int       // 用户ID
	Floor     string    // 楼层
	Content   string    // 评论内容
	Like      int       // 赞
	Unlike    int       // 踩
	IsCheck   int       // 评论是否已审核:0-未审核;1-已审核
	IsDelete  int       // 评论是否已被删除:0-未删除;1-已删除
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 修改时间
}
