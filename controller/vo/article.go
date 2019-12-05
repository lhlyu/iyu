package vo

type ArticleVo struct {
	Id         int    `json:"id"`                          // 文章ID
	UserId     int    `json:"-"`                           // 用户ID
	Wraper     string `json:"wraper"`                      // 头背景
	Title      string `json:"title" validate:"required"`   // 标题
	Content    string `json:"content" validate:"required"` // 内容
	IsTop      int    `json:"isTop"`                       // 是否置顶:1-不置顶;2-置顶
	CategoryId int    `json:"categoryId"`                  // 分类ID
	NailId     int    `json:"nailId"`                      // 钉子ID
	TagArr     []int  `json:"tagArr"`                      // 标签ID组
	Kind       int    `json:"kind"`                        // 文章类型
	IsDelete   int    `json:"isDelete"`                    // 是否已删除:1-未删除;2-已删除
}

type ArticleParam struct {
	PageNum    int    `json:"pageNum" validate:"required,gt=0"`
	PageSize   int    `json:"pageSize" validate:"required,gt=0"`
	CategoryId int    `json:"categoryId"`
	TagId      int    `json:"tagId"`
	Kind       int    `json:"kind"` // 文章类型
	KeyWord    string `json:"keyWord"`
	IsDelete   int    `json:"-"`
}
