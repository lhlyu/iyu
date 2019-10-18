package test

import (
    "fmt"
    "github.com/lhlyu/iyu/common"
    "github.com/lhlyu/iyu/module"
    "testing"
)

func init(){
    module.Register(module.CfgModule,module.DbModule,module.RedisModule)
    module.Init()
}

func TestConfig(t *testing.T){
    fmt.Println(common.Cfg.GetString("version"))
}
