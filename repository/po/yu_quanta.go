package po

import "time"

type YuQuanta struct {
	Id        int       `db:"id"`         // 主键
	Key       string    `db:"key"`        // key值
	Value     string    `db:"value"`      // value值
	Desc      string    `db:"desc"`       // 描述
	IsEnable  int       `db:"is_enable"`  // 是否启用:1-启用;2-不启用
	CreatedAt time.Time `db:"created_at"` // 创建时间
	UpdatedAt time.Time `db:"updated_at"` // 更新时间
}
