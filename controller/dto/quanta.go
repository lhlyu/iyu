package dto

type QuantaDto struct {
	PageNum  int    `json:"pageNum" validate:"required,gt=0"`  // 当前页码
	PageSize int    `json:"pageSize" validate:"required,gt=0"` // 每页记录条数
	Id       int    `json:"id"`
	IsEnable int    `json:"isEnable"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}
