package model

import "time"

type YuNail struct {
	Id        int       `db:"id"`         // 钉子ID
	Name      string    `db:"name"`       // 钉子名字
	Color     string    `db:"color"`      // 钉子颜色
	IsDelete  int       `db:"is_delete"`  // 是否删除: 0-未删除；1-已删除
	CreatedAt time.Time `db:"created_at"` // 创建时间
	UpdatedAt time.Time `db:"updated_at"` // 修改时间
}
