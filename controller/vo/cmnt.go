package vo

type CmntVo struct {
	Id        int    `json:"id"`        // 主键ID
	ArticleId int    `json:"articleId"` // 文章ID
	UserId    int    `json:"userId"`    // 用户ID
	Floor     string `json:"floor"`     // 楼层
	Content   string `json:"content"`   // 评论内容
	IsCheck   int    `json:"-"`         // 评论是否已审核:1-未审核;2-已审核
	IsDelete  int    `json:"-"`         // 评论是否已删除:1-未删除;2-已删除
	PageNum   int    `json:"pageNum"`
	PageSize  int    `json:"pageSize"`
	IsAdmin   bool   `json:"-"`
}
