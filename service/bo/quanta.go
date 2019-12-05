package bo

type Quanta struct {
	Id       int    `json:"id"`
	Key      string `json:"key"`
	Value    string `json:"value"`
	Desc     string `json:"desc"`
	IsEnable int    `json:"isEnable"`
}

type QuantaSort []*Quanta

func (s QuantaSort) Len() int {
	return len(s)
}
func (s QuantaSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s QuantaSort) Less(i, j int) bool {
	return s[i].Key < s[j].Key
}

type QuantaSortByIsEnable struct {
	QuantaSort
}

func (p QuantaSortByIsEnable) Less(i, j int) bool {
	return p.QuantaSort[i].IsEnable < p.QuantaSort[j].IsEnable
}

type QuantaSortByKey struct {
	QuantaSort
}

func (p QuantaSortByKey) Less(i, j int) bool {
	return p.QuantaSort[i].Key < p.QuantaSort[j].Key
}
