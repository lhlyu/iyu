package module

import (
	"github.com/go-gomail/gomail"
	"github.com/lhlyu/iyu/common"
)

type email struct {
}

func (email) SetUp() {
	m := gomail.NewMessage()
	m.SetHeader("From", common.Cfg.GetString("email.from"))
	m.SetHeader("To", common.Cfg.GetString("email.to"))
	m.SetHeader("Subject", common.Cfg.GetString("email.subject"))
	common.Email = common.NewYuEmail(m)
}

// 邮箱模块，主要是别人评论邮箱提示
var EmailModule = email{}
