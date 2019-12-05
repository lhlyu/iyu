package bo

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
