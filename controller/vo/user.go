package vo

type UserParam struct {
	PageNum  int    `json:"pageNum" validate:"required,gt=0"`
	PageSize int    `json:"pageSize" validate:"required,gt=0"`
	KeyWord  string `json:"name"`
}
