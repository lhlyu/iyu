package po

import "time"

type Article struct {
	Id         int       `json:"id"`                     // 主键
	Code       string    `json:"code"`                   // 文章code
	PlateId    int       `json:"plateId"`                // 所属板块ID
	IsTop      int       `json:"isTop" gorm:"default:1"` // 是否置顶:1-否;2-是
	CategoryId int       `json:"categoryId"`             // 分类ID
	Labels     string    `json:"labels"`                 // 标签
	Title      string    `json:"title"`                  // 标题
	Summary    string    `json:"summary"`                // 摘要
	Content    string    `json:"content"`                // 内容
	Cover      string    `json:"cover"`                  // 图片
	State      int       `json:"state" gorm:"default:1"` // 状态:1-正常;2-关闭
	CreatedAt  time.Time `json:"createdAt"`              // 创建时间
	UpdatedAt  time.Time `json:"updatedAt"`              // 修改时间
}
