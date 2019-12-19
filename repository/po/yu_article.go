package po

import "time"

type YuArticle struct {
	Id         int       `json:"id"          db:"id"`          // 文章ID
	Code       string    `json:"code"        db:"code"`        // 文章唯一码
	UserId     int       `json:"userId"      db:"user_id"`     // 用户ID
	Wrapper    string    `json:"wrapper"     db:"wrapper"`     // 头背景
	Title      string    `json:"title"       db:"title"`       // 标题
	Summary    string    `json:"summary"     db:"summary"`     // 摘要/概要
	Content    string    `json:"content"     db:"content"`     // 内容
	IsTop      int       `json:"isTop"       db:"is_top"`      // 是否置顶:1-不置顶;2-置顶
	CategoryId int       `json:"categoryId"  db:"category_id"` // 分类ID
	Kind       int       `json:"kind"        db:"kind"`        // 文章类型:1-普通文章;
	SortNum    int       `json:"sortNum"     db:"sort_num"`    // 排序
	CmntStatus int       `json:"cmntStatus"  db:"cmnt_status"` // 是否开放评论:1-开放;2-不开放
	IsDelete   int       `json:"isDelete"    db:"is_delete"`   // 是否已删除:1-未删除;2-已删除
	Like       int       `json:"like"        db:"like"`        // 赞
	CmntNum    int       `json:"cmntNum"     db:"cmnt_num"`    // 评论量
	Fire       int       `json:"fire"        db:"fire"`        // 浏览量
	CreatedAt  time.Time `json:"createdAt"   db:"created_at"`  // 创建时间
	UpdatedAt  time.Time `json:"updatedAt"   db:"updated_at"`  // 更新时间
}
