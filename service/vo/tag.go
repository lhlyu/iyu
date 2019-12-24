package vo

type TagVo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	IsDelete  int    `json:"isDelete"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}
