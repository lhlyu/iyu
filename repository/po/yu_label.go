package model

import "time"

type YuLabel struct {
	Id        int       // 标签ID
	Name      string    // 标签名字
	IsDelete  int       // 是否删除: 0-未删除；1-已删除
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 修改时间
}
