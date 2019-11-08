package model

type YuUser struct {
	Id        int    `db:"id"`         // 用户id
	ThirdId   int    `db:"third_id"`   // 第三方登录返回的ID
	IsAdmin   int    `db:"is_admin"`   // 是否是管理员:1-普通;2-观察者;9-管理员
	From      int    `db:"from"`       // 来源:1-github;2-gitee
	Status    int    `db:"status"`     // 用户状态:1-正常;2-已删除;3-黑名单
	AvatarUrl string `db:"avatar_url"` // 用户头像
	UserUrl   string `db:"user_url"`   // 用户地址
	UserName  string `db:"user_name"`  // 用户名字
	Bio       string `db:"bio"`        // 个性签名
	Ip        string `db:"ip"`         // ip
	CreatedAt int    `db:"created_at"` // 创建时间
	UpdatedAt int    `db:"updated_at"` // 更新时间
}
