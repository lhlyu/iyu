package common

type base struct {
	traceId string
}

func (b *base) SetTraceId(traceId string) {
	if b == nil {
		return
	}
	b.traceId = traceId
}

func (b *base) GetTraceId() string {
	if b == nil {
		return ""
	}
	return b.traceId
}

// 基础dao
type BaseDao struct {
	base
}

func (s *BaseDao) Error(err error, param ...interface{}) bool {
	if err == nil {
		return false
	}
	Ylog.Log(3, "error", s.traceId, "repository", err.Error(), param)
	return true
}

func (s *BaseDao) Info(param ...interface{}) {
	Ylog.Log(3, "info", s.traceId, "repository", param...)
}

// 基础服务
type BaseService struct {
	base
}

func (s *BaseService) Error(err error, param ...interface{}) bool {
	if err == nil {
		return false
	}
	Ylog.Log(3, "error", s.traceId, "service", err.Error(), param)
	return true
}

func (s *BaseService) Info(param ...interface{}) {
	Ylog.Log(3, "info", s.traceId, "service", param...)
}

// 基础缓存
type BaseCache struct {
	base
}

func (s *BaseCache) Error(err error, param ...interface{}) bool {
	if err == nil {
		return false
	}
	Ylog.Log(3, "error", s.traceId, "cache", err.Error(), param)
	return true
}

func (s *BaseCache) Info(param ...interface{}) {
	Ylog.Log(3, "info", s.traceId, "cache", param...)
}

// 基础控制器
type BaseController struct {
	base
}

func (s BaseController) Error(traceId string, err error, param ...interface{}) bool {
	if err == nil {
		return false
	}
	Ylog.Log(4, "error", traceId, "controller", err.Error(), param)
	return true
}

func (s BaseController) Info(traceId string, param ...interface{}) {
	Ylog.Log(4, "info", traceId, "controller", param...)
}

type MSF = map[string]interface{}
type MSS = map[string]string
type MSI = map[string]int
type MIF = map[int]interface{}
type MIS = map[int]string
type MII = map[int]int
