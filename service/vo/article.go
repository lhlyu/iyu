package vo

// 文章
type Article struct {
	Category *Category `json:"category,omitempty"` // 分类

	Id           uint   `json:"id,omitempty"`           // 自增主键
	Code         string `json:"code,omitempty"`         // 唯一码
	Title        string `json:"title,omitempty"`        // 标题
	Toc          string `json:"toc,omitempty"`          // 目录
	Summary      string `json:"summary,omitempty"`      // 摘要
	Content      string `json:"content,omitempty"`      // 内容
	Cover        string `json:"cover,omitempty"`        // 封面
	Labels       string `json:"labels,omitempty"`       // 标签
	Kind         string `json:"kind,omitempty"`         // 类型:normal(普通),gist(灵感),self(自述)
	State        string `json:"state,omitempty"`        // 状态:draft(草稿),private(私密发布),publish(开放发布),dustbin(垃圾箱)
	Password     string `json:"password,omitempty"`     // 密码,配合私密发布
	CommentState string `json:"commentState,omitempty"` // 评论状态:open(开放),close(关闭),owner(仅所有者)
	Remake       string `json:"remake,omitempty"`       // 备注
	Sort         uint   `json:"sort,omitempty"`         // 排序,降序
	CommentCount uint   `json:"commentCount,omitempty"` // 评论数量
	ViewCount    uint   `json:"viewCount,omitempty"`    // 访问数量
	GoodCount    uint   `json:"goodCount,omitempty"`    // 点赞数量
	BadCount     uint   `json:"badCount,omitempty"`     // 踩数量
	CreatedAt    int64  `json:"createdAt,omitempty"`    // 创建时间
	UpdatedAt    int64  `json:"updatedAt,omitempty"`    // 更新时间
}
