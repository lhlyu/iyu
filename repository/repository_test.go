package repository

import (
	"fmt"
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
	//d.AddNailOne("置顶", "#0000ff")
	a := d.convertToInterface([]int{1, 2, 3, 4, 5})
	fmt.Println(a...)
}

func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

	}
}
