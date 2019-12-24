package po

import "time"

type YuArticleTag struct {
	Id        int       `json:"id"         db:"id"`
	ArticleId int       `json:"articleId"  db:"article_id"`
	TagId     int       `json:"tagId"      db:"tag_id"`
	CreatedAt time.Time `json:"createdAt"  db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt"  db:"updated_at"`
}
