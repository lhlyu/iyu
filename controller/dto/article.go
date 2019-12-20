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
