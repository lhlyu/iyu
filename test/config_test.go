package test

import (
	"fmt"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/module"
	"github.com/lhlyu/iyu/util"
	"testing"
)

func init() {
	module.Register(module.CfgModule, module.RedisModule)
	module.Init()
}

func TestConfig(t *testing.T) {
	fmt.Println(common.Cfg.GetString("version"))
}

type ZZ struct {
	B string
	C int
}

func TestRedis(t *testing.T) {
	common.Redis.HSet("iyu", "1", util.ObjToJsonStr(&ZZ{"tom", 1}))
	common.Redis.HSet("iyu", "2", util.ObjToJsonStr(&ZZ{"tom", 2}))
	common.Redis.HSet("iyu", "3", util.ObjToJsonStr(&ZZ{"tom", 3}))
	common.Redis.HSet("iyu", "4", util.ObjToJsonStr(&ZZ{"tom", 4}))
	common.Redis.HSet("iyu", "5", util.ObjToJsonStr(&ZZ{"tom", 5}))
	common.Redis.HSet("iyu", "6", util.ObjToJsonStr(&ZZ{"tom", 6}))
	common.Redis.HSet("iyu", "6", util.ObjToJsonStr(&ZZ{"tom", 7}))
	m := common.Redis.HMGet("iyu", "2", "3", "6", "5").Val()
	for _, v := range m {
		fmt.Println(v)
	}

}
