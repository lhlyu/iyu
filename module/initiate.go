package module

import "log"

// 启动时执行
type initiate struct {
}

func (initiate) seq() int {
	return 1 << 3
}

func (initiate) SetUp() {
	log.Println("init initiate module ->")
	// 初始化数据
}

var InitiateModule = initiate{}
