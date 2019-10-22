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

func TestSy(t *testing.T) {
	if common.DB == nil {
		return
	}
	d := NewDao()

	fmt.Println(d.Sy().GetErrCode().IsSuccess())
}

func TestRepository(t *testing.T) {
	if common.DB == nil {
		return
	}
	d := NewDao()
	//d.AddNailOne("置顶", "#0000ff")
	a := d.ConvertToInterface([]int{1, 2, 3, 4, 5})
	fmt.Println(a...)
}

func BenchmarkSprintf(b *testing.B) {
	d := NewDao()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.IntConvertToInterface([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	}
}
