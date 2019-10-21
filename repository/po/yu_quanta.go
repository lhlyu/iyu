package model

import "time"

type YuQuanta struct {
	Id        int       `json:"id"`        // 主键
	Key       string    `json:"key"`       // key值
	Value     string    `json:"value"`     // value值
	Desc      string    `json:"desc"`      // 描述
	IsEnable  int       `json:"isEnable"`  // 是否启用:0-启用;1-不启用
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 更新时间
}
