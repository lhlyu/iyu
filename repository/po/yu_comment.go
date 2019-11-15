package model

type YuComment struct {
	Id        int    `db:"id"`         // 主键ID
	ArticleId int    `db:"article_id"` // 文章ID
	UserId    int    `db:"user_id"`    // 用户ID
	Floor     string `db:"floor"`      // 楼层
	Content   string `db:"content"`    // 评论内容
	Like      int    `db:"like"`       // 赞
	Unlike    int    `db:"unlike"`     // 踩
	IsCheck   int    `db:"is_check"`   // 评论是否已审核:1-未审核;2-已审核
	IsDelete  int    `db:"is_delete"`  // 评论是否已删除:1-未删除;2-已删除
	CreatedAt int    `db:"created_at"` // 创建时间
	UpdatedAt int    `db:"updated_at"` // 修改时间
}
