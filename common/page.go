package common

import "math"

type Page struct {
	PageNum     int  `json:"pageNum" validate:"required,gt=0"`  // 当前页码
	PageSize    int  `json:"pageSize" validate:"required,gt=0"` // 每页记录条数
	Total       int  `json:"total"`                             // 记录总量
	PageMax     int  `json:"pageMax"`                           // 最大页码
	PrePage     int  `json:"prePage"`                           // 上一页码,如果没有为 0
	NextPage    int  `json:"nextPage"`                          // 下一页码,如果没有为 0
	HasPrePage  bool `json:"hasPrePage"`                        // 是否有上一页
	HasNextPage bool `json:"hasNextPage"`                       // 是否有下一页
	Remainder   int  `json:"remainder"`                         // 剩余数据量
	StartRow    int  `json:"-"`                                 // 记录开始行
	StopRow     int  `json:"-"`                                 // 记录结束行
}

func NewPage(pageNum, pageSize int) *Page {
	return &Page{
		PageNum:  pageNum,
		PageSize: pageSize,
	}
}

func (p *Page) SetTotal(total int) {
	if p == nil {
		return
	}
	p.Total = total
	if p.PageSize <= 0 {
		return
	}
	p.PageMax = int(math.Ceil(float64(p.Total) / float64(p.PageSize)))
	p.StartRow = (p.PageNum - 1) * p.PageSize
	p.StopRow = p.StartRow + p.PageSize
	p.PrePage = p.PageNum - 1
	p.NextPage = p.PageNum + 1
	if p.PrePage <= 0 {
		p.PrePage = 0
		p.HasPrePage = false
	} else {
		p.HasPrePage = true
	}
	if p.NextPage >= p.PageMax {
		p.NextPage = 0
		p.HasNextPage = false
	} else {
		p.HasNextPage = true
	}
	if p.StartRow > p.Total {
		p.StartRow = p.Total
	}
	if p.StopRow > p.Total {
		p.StopRow = p.Total
	}
	p.Remainder = p.Total - p.StopRow
}
