package article_service

import (
	"github.com/lhlyu/iyu/module"
	"github.com/lhlyu/yutil"
	"testing"
)

func init() {
	module.Register(module.CfgModule, // 读取配置 <必须>
		module.LgModule, // 日志
		module.DbModule, // 连接数据库
		module.InitiateModule,
		module.RedisModule) // redis
	module.Init()
}

func TestService_GetArticleById(t *testing.T) {
	svc := NewService("test")
	result := svc.GetArticleById(1)
	t.Log(yutil.JsonObjToStr(result))
}
