package model

type YuNail struct {
	Id        int    `db:"id"`         // 钉子ID
	Name      string `db:"name"`       // 钉子名字
	Color     string `db:"color"`      // 钉子颜色
	IsDelete  int    `db:"is_delete"`  // 是否已删除:1-未删除;2-已删除
	CreatedAt int    `db:"created_at"` // 创建时间
	UpdatedAt int    `db:"updated_at"` // 修改时间
}
