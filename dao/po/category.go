package po

import (
	"time"
)

type Category struct {
	Id        uint      `db:"id"`         // 自增主键
	Name      string    `db:"name"`       // 名字
	Count     uint      `db:"count"`      // 包含文章数量
	CreatedAt time.Time `db:"created_at"` // 创建时间
	UpdatedAt time.Time `db:"updated_at"` // 更新时间
}
