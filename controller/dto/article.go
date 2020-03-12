package dto

// 文章
type Article struct {
	Id         uint   `json:"id"`         // 主键
	Code       string `json:"code"`       // 文章code
	PlateId    uint   `json:"plateId"`    // 所属板块ID
	IsTop      uint   `json:"isTop"`      // 是否置顶:1-否;2-是
	CategoryId uint   `json:"categoryId"` // 分类ID
	Color      string `json:"color"`      // 颜色
	Labels     string `json:"labels"`     // 标签
	Title      string `json:"title"`      // 标题
	Summary    string `json:"summary"`    // 摘要
	Content    string `json:"content"`    // 内容
	Cover      string `json:"cover"`      // 图片
	State      uint   `json:"state"`      // 状态:1-正常;2-关闭
}
