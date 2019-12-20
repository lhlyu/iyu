package vo

type ArticleVo struct {
	ArticleData
	Id       int `json:"id"`
	Kind     int `json:"kind"`
	IsDelete int `json:"isDelete"`
	SortNum  int `json:"sortNum"`
}
type ArticleAuthor struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
type ArticleCategory struct {
	ID    int    `json:"id"`
	Color string `json:"color"`
	Name  string `json:"name"`
}
type ArticleTags struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type ArticleStat struct {
	Fire    int `json:"fire"`
	Like    int `json:"like"`
	CmntNum int `json:"cmntNum"`
}

type ArticleData struct {
	Code       string           `json:"code"`
	Wrapper    string           `json:"wrapper"`
	Title      string           `json:"title"`
	Summmary   string           `json:"summmary"`
	Content    string           `json:"content"`
	IsTop      int              `json:"isTop"`
	CmntStatus int              `json:"cmntStatus"`
	CreatedAt  int64            `json:"createdAt"`
	UpdateAt   int64            `json:"updatedAt"`
	Author     *ArticleAuthor   `json:"author:"`
	Category   *ArticleCategory `json:"category"`
	Tags       []*ArticleTags   `json:"tags"`
	Stat       *ArticleStat     `json:"stat"`
}
