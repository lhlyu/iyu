package model

import "time"

type YuUserInfo struct {
	Id        int
	UserId    int    // 用户ID
	AvatarUrl string // 用户头像
	UserUrl   string // 用户地址
	UserName  string // 用户名字
	Bio       string // 个性签名
	Ip        string // ip
	CreatedAt time.Time
	UpdatedAt time.Time
}
