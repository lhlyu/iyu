package model

import "time"

type YuNail struct {
	Id        int       // 钉子ID
	Name      string    // 钉子名字
	Color     string    // 钉子颜色
	IsDelete  int       // 是否删除: 0-未删除；1-已删除
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 修改时间
}
