package cache

import "github.com/lhlyu/iyu/common"

type Cache struct {
	common.BaseCache
}

func NewCache(traceId string) *Cache {
	che := &Cache{}
	che.SetTraceId(traceId)
	return che
}

/**
常用元素: 网站配置 分类 标签
type: hash
survival: week
named:
  iyu:quanta   key   value
  iyu:category id    value
  iyu:tag      id    value

用户:
type: hash
survival: day
named:
    iyu:user  userId   value

type: set
survival: day

named:
    iyu:global:visit  userId
    iyu:article:visit:{id}  userId
    iyu:article:like:{id}   userId
    iyu:cmnt:like:{id}      userId
    iyu:reply:like:{id}     userId

article

type: hash
survival: day

named:
    iyu:article  id  value


type : string  nx
token:
    iyu:token:{userId}  value

future:
网站访问量 文章访问量
*/
