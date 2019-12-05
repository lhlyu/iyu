package bo

type Category struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	IsDelete int    `json:"isDelete"`
}

type CategorySort []*Category

func (s CategorySort) Len() int {
	return len(s)
}
func (s CategorySort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s CategorySort) Less(i, j int) bool {
	return s[i].IsDelete < s[j].IsDelete
}
