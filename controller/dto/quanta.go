package dto

import "github.com/lhlyu/iyu/common"

type QuantaDto struct {
	*common.Page
	Id       int    `json:"id"`
	IsEnable int    `json:"isEnable"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}
