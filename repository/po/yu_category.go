package po

import "time"

type YuCategory struct {
    Id         int        `db:"id"`           // 分类ID
    Name       string     `db:"name"`         // 分类名字
    IsDelete   int        `db:"is_delete"`    // 是否已删除:1-未删除;2-已删除
    CreatedAt  time.Time  `db:"created_at"`   // 创建时间
    UpdatedAt  time.Time  `db:"updated_at"`   // 更新时间
}

