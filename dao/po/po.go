package po

// 这个不属于数据库的任何表，仅仅用于 定义枚举值

// 通用 state
const (
	STATE_ONE = iota + 1
	STATE_TEO
)

// record kind
const (
	RECORD_KIND_SYSTEM = iota + 1
	RECORD_KIND_ERROR
	RECORD_KIND_VISIT
)
