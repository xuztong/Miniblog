package common

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewPageRequest() *PageRequest {
	return &PageRequest{
		PageSize:   10,
		PageNumber: 1,
	}
}

func NewPageRequestFromGinCtx(c *gin.Context) *PageRequest {
	p := NewPageRequest()
	pnStr := c.Query("page_number")
	psStr := c.Query("page_size")
	if pnStr != "" {
		pn, _ := strconv.ParseInt(pnStr, 10, 64)
		if pn != 0 {
			p.PageNumber = int(pn)
		}
	}
	if psStr != "" {
		ps, _ := strconv.ParseInt(psStr, 10, 64)
		if ps != 0 {
			p.PageSize = int(ps)
		}
	}
	return p
}

type PageRequest struct {
	PageSize   int `json:"page_size"`
	PageNumber int `json:"page_number"`
}

func (p *PageRequest) Offset() int {
	return (p.PageNumber - 1) * p.PageSize
}
