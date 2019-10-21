package model

type YuArticleLabel struct {
	Id        int
	ArticleId int
	LabelId   int
	IsDelete  int // 是否删除: 0-未删除；1-已删除
}
