package vo

type UserVo struct {
	Id        int    `json:"id"`
	ThirdId   int    `json:"thirdId"`
	Role      int    `json:"role"`
	From      int    `json:"from"`
	Status    int    `json:"status"`
	AvatarUrl string `json:"avatarUrl"`
	UserUrl   string `json:"userUrl"`
	UserName  string `json:"userName"`
	Bio       string `json:"bio"`
	Sign      string `json:"sign"`
	Ip        string `json:"ip"`
	LastLogin int64  `json:"lastLogin"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}
