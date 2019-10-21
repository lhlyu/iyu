package model

import "time"

type YuCategory struct {
	Id        int       // 分类ID
	Name      string    // 分类名字
	IsDelete  int       // 是否删除:0-未删除;1-已删除
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}
