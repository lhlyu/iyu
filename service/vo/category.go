package vo

type Category struct {
	Id    uint   `json:"id"`    // 自增主键
	Name  string `json:"name"`  // 名字
	Count uint   `json:"count"` // 包含文章数量
}
