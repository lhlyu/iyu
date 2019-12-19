package po

import "time"

type YuCmnt struct {
	Id        int       `json:"id"         db:"id"`
	ArticleId int       `json:"articleId"  db:"article_id"` // 文章ID
	SortNum   int       `json:"sortNum"    db:"sort_num"`   // 排序
	Floor     int       `json:"floor"      db:"floor"`      // 楼层
	Like      int       `json:"like"       db:"like"`       // 赞
	Content   string    `json:"content"    db:"content"`    // 内容
	IsDelete  int       `json:"isDelete"   db:"is_delete"`  // 删除:1-未删除;2-已删除
	IsCheck   int       `json:"isCheck"    db:"is_check"`   // 审核:1-未审核;2-已审核
	ReplyNum  int       `json:"replyNum"   db:"reply_num"`  // 回复数量
	UserId    int       `json:"userId"     db:"user_id"`    // 用户ID
	CreatedAt time.Time `json:"createdAt"  db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt"  db:"updated_at"`
}
