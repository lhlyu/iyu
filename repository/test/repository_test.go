package test

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/module"
	"github.com/lhlyu/iyu/repository"
	"testing"
)

func init() {
	module.Register(module.CfgModule, module.DbModule, module.LgModule)
	module.Init()
}

func TestQuery(t *testing.T) {
	if common.DB == nil {
		return
	}
	d := repository.NewDao()
	user, err := d.GetUserById(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(user)
}

func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

	}
}
