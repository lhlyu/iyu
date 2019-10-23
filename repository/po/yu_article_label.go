package model

type YuArticleLabel struct {
	Id        int `db:"id"`
	ArticleId int `db:"article_id"`
	LabelId   int `db:"label_id"`
	IsDelete  int `db:"is_delete"` // 是否已删除:1-未删除;2-已删除
}
