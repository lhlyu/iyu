package module

import "github.com/lhlyu/iyu/common"

type lg struct {
}

func (lg) seq() int {
	return 1 << 0
}

func (lg) SetUp() {
	common.Ylog = common.NewYlog(common.Cfg.GetString("log.level"),
		common.Cfg.GetString("log.timeFormat"),
		common.Cfg.GetString("log.outFile"),
		common.Cfg.GetString("log.outWay"))
}

var LgModule = lg{}
