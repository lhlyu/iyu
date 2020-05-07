package service

import "github.com/lhlyu/iyu/trace"

type BaseService struct {
	trace.BaseTracker
}

func NewBaseService(tracker trace.ITracker) BaseService {
	return BaseService{
		BaseTracker: trace.NewBaseTracker(tracker),
	}
}
