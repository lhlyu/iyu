package po

import "time"

type YuTag struct {
	Id        int       `json:"id"         db:"id"`         // 标签ID
	Name      string    `json:"name"       db:"name"`       // 标签名字
	IsDelete  int       `json:"isDelete"   db:"is_delete"`  // 是否已删除:1-未删除;2-已删除
	CreatedAt time.Time `json:"createdAt"  db:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"  db:"updated_at"` // 修改时间
}
