package po

import "time"

type YuQuanta struct {
	Id        int       `json:"id"         db:"id"`         // 主键
	Key       string    `json:"key"        db:"key"`        // key值
	Value     string    `json:"value"      db:"value"`      // value值
	Desc      string    `json:"desc"       db:"desc"`       // 描述
	IsEnable  int       `json:"isEnable"   db:"is_enable"`  // 是否启用:1-启用;2-不启用
	CreatedAt time.Time `json:"createdAt"  db:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"  db:"updated_at"` // 更新时间
}
