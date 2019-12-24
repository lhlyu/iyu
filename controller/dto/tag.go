package dto

import "github.com/lhlyu/iyu/common"

type TagDto struct {
	*common.Page
	Id       int    `json:"id"`
	IsDelete int    `json:"isDelete"`
	Name     string `json:"name"`
}
