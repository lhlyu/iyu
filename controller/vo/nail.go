package vo

type NailVo struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Color  string `json:"color"`
	Status int    `json:"status"`
	Real   int    `json:"real"`
}
