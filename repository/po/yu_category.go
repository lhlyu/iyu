package po

import "time"

type YuCategory struct {
	Id        int       `json:"id"         db:"id"`         // 分类ID
	Name      string    `json:"name"       db:"name"`       // 分类名字
	Color     string    `json:"color"      db:"color"`      // 颜色
	IsDelete  int       `json:"isDelete"   db:"is_delete"`  // 是否已删除:1-未删除;2-已删除
	CreatedAt time.Time `json:"createdAt"  db:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"  db:"updated_at"` // 更新时间
}
