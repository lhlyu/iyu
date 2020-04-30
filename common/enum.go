package common

// kind type
const (
	KIND_NORMAL = "normal" // 普通
	KIND_GIST   = "gist"   // 灵感
	KIND_SELF   = "self"   // 自述
)

// state
const (
	STATE_DRAFT   = "draft"   // 草稿
	STATE_PRIVATE = "private" // 私密发布
	STATE_PUBLISH = "publish" // 开放发布
	STATE_DUSTBIN = "dustbin" // 垃圾箱
)

// comment status
const (
	COMMENT_OPEN  = "open"  // 开放
	COMMENT_CLOSE = "close" // 关闭
	COMMENT_OWNER = "owner" // 仅所有者
)
