package test

import (
    "fmt"
    "github.com/lhlyu/iyu/common"
    "github.com/lhlyu/iyu/module"
    "testing"
    "time"
)

func init(){
    module.Register(module.CfgModule,module.RedisModule)
    module.Init()
}

func TestConfig(t *testing.T){
    fmt.Println(common.Cfg.GetString("version"))
    common.Redis.Set("LHLYU","123",time.Hour)
}
