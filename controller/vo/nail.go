package vo

type NailVo struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Color    string `json:"color"`
	Real     int    `json:"real"`
	IsDelete int    `json:"isDelete"`
}
