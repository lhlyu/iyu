package po

import "time"

type YuRecord struct {
	Id           int       `db:"id"`
	UserId       int       `db:"user_id"`
	BusinessId   int       `db:"business_id"`   // 目标Id
	BusinessKind int       `db:"business_kind"` // 目标类型:1-文章;2-评论;3-回复
	Action       int       `db:"action"`        // 动作:1-浏览;2-评论;3-赞;4-踩
	Ip           string    `db:"ip"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type Stat struct {
	Action int `db:"action"`
	Number int `db:"number"`
}
