package model

import "time"

type YuUser struct {
	Id        int       // 用户id
	ThirdId   int       // 第三方登录返回的ID
	IsAdmin   int       // 是否是管理员:0-普通;1-观察者;9-管理员
	From      int       // 来源:1-github;2-gitee
	Status    int       // 用户状态:0-正常;1-已删除;2-黑名单
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}
