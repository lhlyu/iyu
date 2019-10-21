package model

import "time"

type YuUserInfo struct {
	Id        int       `json:"id"`
	UserId    int       `json:"userId"`    // 用户ID
	AvatarUrl string    `json:"avatarUrl"` // 用户头像
	UserUrl   string    `json:"userUrl"`   // 用户地址
	UserName  string    `json:"userName"`  // 用户名字
	Bio       string    `json:"bio"`       // 个性签名
	Ip        string    `json:"ip"`        // ip
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
