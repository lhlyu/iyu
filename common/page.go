package common

import "math"

type Page struct {
	PageNum     int  `json:"pageNum" validate:"required,gt=0"`
	PageSize    int  `json:"pageNum" validate:"required,gt=0"`
	Total       int  `json:"total"`
	PageMax     int  `json:"pageMax"`
	PrePage     int  `json:"prePage"`
	NextPage    int  `json:"nextPage"`
	hasPrePage  bool `json:"hasPrePage"`
	hasNextPage bool `json:"hasNextPage"`
	StartRow    int  `json:"-"`
}

func NewPage(pageNum, pageSize int) *Page {
	return &Page{
		PageNum:  pageNum,
		PageSize: pageSize,
	}
}

func (p *Page) SetTotal(total int) {
	p.Total = total
	if p.PageSize <= 0 {
		return
	}
	p.PageMax = int(math.Ceil(float64(p.Total) / float64(p.PageSize)))
	p.StartRow = (p.PageNum - 1) * p.PageSize
	p.PrePage = p.PageNum - 1
	p.NextPage = p.PageNum + 1
	if p.PrePage <= 0 {
		p.hasPrePage = false
	} else {
		p.hasPrePage = true
	}
	if p.NextPage >= p.PageMax {
		p.hasNextPage = false
	} else {
		p.hasNextPage = true
	}
}
