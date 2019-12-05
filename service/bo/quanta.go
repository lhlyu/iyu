package bo

type Quanta struct {
	Id       int    `json:"id"`
	Key      string `json:"key"`
	Value    string `json:"value"`
	Desc     string `json:"desc"`
	IsEnable int    `json:"isEnable"`
}
