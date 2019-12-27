package dto

type CategoryDto struct {
	PageNum  int    `json:"pageNum" validate:"required,gt=0"`  // 当前页码
	PageSize int    `json:"pageSize" validate:"required,gt=0"` // 每页记录条数
	Id       int    `json:"id"`
	IsDelete int    `json:"isDelete"`
	Name     string `json:"name"`
	Color    string `json:"color"`
}
