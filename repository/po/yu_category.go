package model

import "time"

type YuCategory struct {
	Id        int       `json:"id"`        // 分类ID
	Name      string    `json:"name"`      // 分类名字
	IsDelete  int       `json:"isDelete"`  // 是否删除:0-未删除;1-已删除
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 更新时间
}
