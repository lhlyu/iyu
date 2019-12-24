package dto

import "github.com/lhlyu/iyu/common"

type ArticleDto struct {
	*common.Page
	Id         int    `json:"id"`
	Code       string `json:"code"`
	KeyWord    string `json:"keyWord"`
	CategoryId int    `json:"categoryId"`
	TagId      int    `json:"tagId"`
	Kind       int    `json:"-"`
	IsDelete   int    `json:"-"`
}

type ArticleEditDto struct {
	UserId     int    `json:"-"`
	Wrapper    string `json:"wrapper"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	Content    string `json:"content"`
	IsTop      int    `json:"isTop"`
	CategoryId int    `json:"categoryId"`
	Kind       int    `json:"kind"`
	SortNum    int    `json:"sortNum"`
	CmntStatus int    `json:"cmntStatus"`
	IsDelete   int    `json:"isDelete"`
	TagIds     []int  `json:"tagIds"`
}
