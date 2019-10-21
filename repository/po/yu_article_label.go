package model

type YuArticleLabel struct {
	Id        int `json:"id"`
	ArticleId int `json:"articleId"`
	LabelId   int `json:"labelId"`
	IsDelete  int `json:"isDelete"` // 是否删除: 0-未删除；1-已删除
}
