package bo

type ArticleData struct {
	ID        int       `json:"id"`
	Kind      int       `json:"kind"`
	Fire      int       `json:"fire"`
	CmntNum   int       `json:"cmntNum"`
	Like      int       `json:"like"`
	UnLike    int       `json:"unlike"`
	CreatedAt int       `json:"createdAt"`
	UpdatedAt int       `json:"updatedAt"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Wraper    string    `json:"wraper"`
	Nail      *Nail     `json:"nail"`
	Category  *Category `json:"category"`
	Tags      []*Tag    `json:"tags"`
}

type Article struct {
	Id        int       `json:"id"`
	Kind      int       `json:"kind"`
	Fire      int       `json:"fire"`
	CmntNum   int       `json:"cmntNum"`
	Like      int       `json:"like"`
	UnLike    int       `json:"unlike"`
	CreatedAt int       `json:"createdAt"`
	UpdatedAt int       `json:"updatedAt"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Wraper    string    `json:"wraper"`
	Nail      *Nail     `json:"nail"`
	Category  *Category `json:"category"`
	Tags      []*Tag    `json:"tags"`
}
