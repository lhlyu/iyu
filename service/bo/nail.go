package bo

type Nail struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Color    string `json:"color"`
	IsDelete int    `json:"isDelete"`
}

type NailSort []*Nail

func (s NailSort) Len() int {
	return len(s)
}
func (s NailSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s NailSort) Less(i, j int) bool {
	return s[i].IsDelete < s[j].IsDelete
}
