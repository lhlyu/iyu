package common

import "strings"

const (
	LV0 = iota // 游客
	LV1        // 普通
	LV2
	LV3
	LV4
	LV5
	LV6
	LV7
	LV8 // 管理员
	LV9 // 系统管理员
)

const PERMISSION = LV8

const (
	X_ID   = "i"
	X_ROLE = "r"
	X_TIME = "t"
	X_USER = "x"
)

type XUser struct {
	Id   int
	Role int
	Ip   string
}

func (x *XUser) GetIp() string {
	return strings.ReplaceAll(x.Ip, ".", ":")
}
