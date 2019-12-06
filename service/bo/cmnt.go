package bo

type Cmnt struct {
	Id         int       `json:"id"`        // 主键ID
	ArticleId  int       `json:"articleId"` // 文章ID
	UserId     int       `json:"userId"`    // 用户ID
	UserName   string    `json:"userName"`
	UserAvatar string    `json:"userAvatar"`
	Floor      string    `json:"floor"`              // 楼层
	Content    string    `json:"content"`            // 评论内容
	IsCheck    int       `json:"isCheck,omitempty"`  // 评论是否已审核:1-未审核;2-已审核
	IsDelete   int       `json:"isDelete,omitempty"` // 评论是否已删除:1-未删除;2-已删除
	CreatedAt  int       `json:"createdAt"`          // 创建时间
	UpdatedAt  int       `json:"updatedAt"`          // 修改时间
	PostList   *PostData `json:"postList,omitempty"`
}
