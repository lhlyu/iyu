package po

import "time"

type YuArticleTag struct {
	Id        int       `db:"id"`
	ArticleId int       `db:"article_id"`
	TagId     int       `db:"tag_id"`
	IsDelete  int       `db:"is_delete"`  // 是否已删除:1-未删除;2-已删除
	CreatedAt time.Time `db:"created_at"` // 创建时间
	UpdatedAt time.Time `db:"updated_at"` // 更新时间
}
