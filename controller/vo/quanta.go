package vo

type QuantaVo struct {
	Id       int    `json:"id"`
	Key      string `json:"key"`
	Value    string `json:"value"`
	Desc     string `json:"desc"`
	IsEnable int    `json:"is_enable"`
	Real     int    `json:"real"`
}
