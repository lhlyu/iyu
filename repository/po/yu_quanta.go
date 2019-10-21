package model

import "time"

type YuQuanta struct {
	Id        int       // 主键
	Key       string    // key值
	Value     string    // value值
	Desc      string    // 描述
	IsEnable  int       // 是否启用:0-启用;1-不启用
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}
