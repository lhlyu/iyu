package repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/module"
	"testing"
)

func init() {
	module.Register(module.CfgModule, module.DbModule)
	module.Init()
}

func TestRepository(t *testing.T) {
	if common.DB == nil {
		return
	}
	d := NewDao()
	d.AddNailOne("置顶", "#0000ff")
}
