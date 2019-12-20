package vo

type QuantaVo struct {
	Id        int    `json:"id"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	Desc      string `json:"desc"`
	IsEnable  int    `json:"isEnable"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}
