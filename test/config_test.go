package test

import (
	"fmt"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/module"
	"testing"
	"time"
)

func init() {
	module.Register(module.CfgModule, module.RedisModule)
	module.Init()
}

func TestConfig(t *testing.T) {
	fmt.Println(common.Cfg.GetString("version"))
}

func TestRedis(t *testing.T) {
	ok := common.Redis.SetNX("IYU", "1", time.Hour*24).Val()
	fmt.Println(ok)
}
