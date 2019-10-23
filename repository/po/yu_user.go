package model

import "time"

type YuUser struct {
	Id        int       `db:"id"`         // 用户id
	ThirdId   int       `db:"third_id"`   // 第三方登录返回的ID
	IsAdmin   int       `db:"is_admin"`   // 是否是管理员:1-普通;2-观察者;9-管理员
	From      int       `db:"from"`       // 来源:1-github;2-gitee
	Status    int       `db:"status"`     // 用户状态:1-正常;2-已删除;3-黑名单
	CreatedAt time.Time `db:"created_at"` // 创建时间
	UpdatedAt time.Time `db:"updated_at"` // 更新时间
}
