package bo

type ArticleParam struct {
	PageNum    int `json:"pageNum" validate:"required,gt=0"`
	PageSize   int `json:"pageSize" validate:"required,gt=0"`
	CategoryId int `json:"categoryId"`
	TagId      int `json:"tagId"`
}
