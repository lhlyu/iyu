package dto

// 文章
type Article struct {
	Id           uint   `json:"id"`           // 自增主键
	Code         string `json:"code"`         // 唯一码
	Title        string `json:"title"`        // 标题
	Toc          string `json:"toc"`          // 目录
	Summary      string `json:"summary"`      // 摘要
	Content      string `json:"content"`      // 内容
	Cover        string `json:"cover"`        // 封面
	Labels       string `json:"labels"`       // 标签
	Kind         string `json:"kind"`         // 类型:normal(普通),gist(灵感),self(自述)
	State        string `json:"state"`        // 状态:draft(草稿),private(私密发布),publish(开放发布),dustbin(垃圾箱)
	Password     string `json:"password"`     // 密码,配合私密发布
	CommentState string `json:"commentState"` // 评论状态:open(开放),close(关闭),owner(仅所有者)
	Remake       string `json:"remake"`       // 备注
	Sort         uint   `json:"sort"`         // 排序,降序
	Category     uint   `json:"category"`     // 分类
	CommentCount uint   `json:"commentCount"` // 评论数量
	ViewCount    uint   `json:"viewCount"`    // 访问数量
	GoodCount    uint   `json:"goodCount"`    // 点赞数量
	BadCount     uint   `json:"badCount"`     // 踩数量
	CreatedAt    int64  `json:"createdAt"`    // 创建时间
	UpdatedAt    int64  `json:"updatedAt"`    // 更新时间
}
