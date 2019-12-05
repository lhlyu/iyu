package bo

type UserData struct {
	Id        int    `db:"id"`       // 用户id
	ThirdId   int    `db:"third_id"` // 第三方登录返回的ID
	Role      int    `db:"role"`
	From      int    `db:"from"`       // 来源:1-github;2-gitee
	Status    int    `db:"status"`     // 用户状态:1-正常;2-已删除;3-黑名单
	AvatarUrl string `db:"avatar_url"` // 用户头像
	UserUrl   string `db:"user_url"`   // 用户地址
	UserName  string `db:"user_name"`  // 用户名字
	Bio       string `db:"bio"`        // 个性签名
	Ip        string `db:"ip"`         // ip
	CreatedAt int    `db:"created_at"` // 创建时间
	UpdatedAt int    `db:"updated_at"` // 更新时间
}

type User struct {
	Id        int    `json:"id"`      // 用户id
	ThirdId   int    `json:"thirdId"` // 第三方登录返回的ID
	Role      int    `json:"role"`
	From      int    `json:"from"`      // 来源:1-github;2-gitee
	Status    int    `json:"status"`    // 用户状态:1-正常;2-已删除;3-黑名单
	AvatarUrl string `json:"avatarUrl"` // 用户头像
	UserUrl   string `json:"userUrl"`   // 用户地址
	UserName  string `json:"userName"`  // 用户名字
	Bio       string `json:"bio"`       // 个性签名
	Ip        string `json:"ip"`        // ip
	CreatedAt int    `json:"createdAt"` // 创建时间
	UpdatedAt int    `json:"updatedAt"` // 更新时间
}

type UserSort []*User

func (s UserSort) Len() int {
	return len(s)
}
func (s UserSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s UserSort) Less(i, j int) bool {
	return s[i].Role > s[j].Role
}

type UserSortByCreatedAt struct {
	UserSort
}

func (p UserSortByCreatedAt) Less(i, j int) bool {
	return p.UserSort[i].CreatedAt > p.UserSort[j].CreatedAt
}
