package dao

import (
	"github.com/lhlyu/iyu/trace"
)

type BaseDao struct {
	trace.BaseTracker
}

func NewBaseDao(tracker trace.ITracker) BaseDao {
	return BaseDao{
		BaseTracker: trace.NewBaseTracker(tracker),
	}
}
