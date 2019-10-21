package model

type YuArticleCatalog struct {
	Id        int
	ArticleId int
	CatalogId int
	IsDelete  int // 是否删除: 0-未删除；1-已删除
}
