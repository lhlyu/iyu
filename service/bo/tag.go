package bo

type Tag struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	IsDelete int    `json:"isDelete"`
}

type TagSort []*Tag

func (s TagSort) Len() int {
	return len(s)
}
func (s TagSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s TagSort) Less(i, j int) bool {
	return s[i].IsDelete < s[j].IsDelete
}
