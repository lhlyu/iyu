package po

import "time"

type YuUser struct {
	Id        int       `json:"id"         db:"id"`         // 用户id
	ThirdId   int       `json:"thirdId"    db:"third_id"`   // 第三方登录返回的ID
	Role      int       `json:"role"       db:"role"`       // 1-游客;2-普通;3-观察者;8-管理员;9-系统管理员
	From      int       `json:"from"       db:"from"`       // 来源:1-github;2-gitee
	Status    int       `json:"status"     db:"status"`     // 用户状态:1-正常;2-已删除;3-黑名单(禁止发言)
	AvatarUrl string    `json:"avatarUrl"  db:"avatar_url"` // 用户头像
	UserUrl   string    `json:"userUrl"    db:"user_url"`   // 用户地址
	UserName  string    `json:"userName"   db:"user_name"`  // 用户名字
	Bio       string    `json:"bio"        db:"bio"`        // 个性签名
	Sign      string    `json:"sign"       db:"sign"`       // 个性图片
	Ip        string    `json:"ip"         db:"ip"`         // ip
	LastLogin time.Time `json:"lastLogin"  db:"last_login"` // 最近访问
	CreatedAt time.Time `json:"createdAt"  db:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"  db:"updated_at"` // 更新时间
}
