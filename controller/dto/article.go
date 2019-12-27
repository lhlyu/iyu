package dto

type ArticleDto struct {
	PageNum    int    `json:"pageNum" validate:"required,gt=0"`  // 当前页码
	PageSize   int    `json:"pageSize" validate:"required,gt=0"` // 每页记录条数
	Id         int    `json:"id"`
	Code       string `json:"code"`
	KeyWord    string `json:"keyWord"`
	CategoryId int    `json:"categoryId"`
	TagId      int    `json:"tagId"`
	Kind       int    `json:"-"`
	IsDelete   int    `json:"-"`
}

type ArticleEditDto struct {
	UserId     int    `json:"-"`
	Wrapper    string `json:"wrapper"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	Content    string `json:"content"`
	IsTop      int    `json:"isTop"`
	CategoryId int    `json:"categoryId"`
	Kind       int    `json:"kind"`
	SortNum    int    `json:"sortNum"`
	CmntStatus int    `json:"cmntStatus"`
	IsDelete   int    `json:"isDelete"`
	TagIds     []int  `json:"tagIds"`
}
