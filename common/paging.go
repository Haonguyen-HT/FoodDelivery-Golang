package common

import "strings"

type Paging struct {
	Page       int    `json:"page" form:"page"`
	Limit      int    `json:"limit" form:"limit"`
	Total      int64  `json:"total" form:"total"`
	Cursor     string `json:"cursor" form:"cursor"`
	NextCursor string `json:"next_cursor" form:"next_cursor"`
}

func (p *Paging) Fulfill() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 10
	}

	p.Cursor = strings.TrimSpace(p.Cursor)
}
