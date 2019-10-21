package model

import "time"

type YuLabel struct {
	Id        int       `json:"id"`        // 标签ID
	Name      string    `json:"name"`      // 标签名字
	IsDelete  int       `json:"isDelete"`  // 是否删除: 0-未删除；1-已删除
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 修改时间
}
