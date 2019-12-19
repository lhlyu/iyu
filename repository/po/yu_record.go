package po

import "time"

type YuRecord struct {
	Id           int       `json:"id"            db:"id"`
	UserId       int       `json:"userId"        db:"user_id"`       // 用户ID
	BusinessId   int       `json:"businessId"    db:"business_id"`   // 目标Id
	Content      string    `json:"content"       db:"content"`       // 内容
	BusinessKind int       `json:"businessKind"  db:"business_kind"` // 1.系统操作;2.错误日志;3.用户登录;4.全站浏览;5.文章浏览;6.文章赞;7.文章踩;8.文章评论;9.评论赞;10.评论踩;11.评论回复;12.回复赞;13.回复踩
	Agent        string    `json:"agent"         db:"agent"`
	Ip           string    `json:"ip"            db:"ip"` // IP地址
	CreatedAt    time.Time `json:"createdAt"     db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt"     db:"updated_at"`
}
