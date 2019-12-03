package vo

type UserParam struct {
	PageNum  int    `json:"pageNum" validate:"required,gt=0"`
	PageSize int    `json:"pageSize" validate:"required,gt=0"`
	KeyWord  string `json:"name"`
	Id       int    `json:"-"`
}

type UserEditParam struct {
	Id        int
	Role      int
	Status    int
	AvatarUrl string
	UserName  string
	UserUrl   string
	Bio       string
	Ip        string
	ThirdId   int
	From      string
}
