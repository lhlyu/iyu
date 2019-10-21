package model

type YuArticleCatalog struct {
	Id        int `json:"id"`
	ArticleId int `json:"articleId"`
	CatalogId int `json:"catalogId"`
	IsDelete  int `json:"isDelete"` // 是否删除: 0-未删除；1-已删除
}
