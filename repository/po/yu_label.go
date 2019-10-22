package model

import "time"

type YuLabel struct {
	Id        int       `db:"id"`         // 标签ID
	Name      string    `db:"name"`       // 标签名字
	IsDelete  int       `db:"is_delete"`  // 是否删除: 0-未删除；1-已删除
	CreatedAt time.Time `db:"created_at"` // 创建时间
	UpdatedAt time.Time `db:"updated_at"` // 修改时间
}
