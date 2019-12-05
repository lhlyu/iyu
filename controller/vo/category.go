package vo

type CategoryVo struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Real     int    `json:"real"`
	IsDelete int    `json:"isDelete"`
}
