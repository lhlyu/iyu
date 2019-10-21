package model

import "time"

type YuNail struct {
	Id        int       `json:"id"`        // 钉子ID
	Name      string    `json:"name"`      // 钉子名字
	Color     string    `json:"color"`     // 钉子颜色
	IsDelete  int       `json:"isDelete"`  // 是否删除: 0-未删除；1-已删除
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 修改时间
}
