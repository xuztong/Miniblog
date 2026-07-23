package blog

import (
	"blog/common"
	"time"
)

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{
		Tags: map[string]string{},
	}
}

type CreateBlogRequest struct {
	Title    string            `json:"title" gorm:"column:title" validate:"required"`
	Author   string            `json:"author" gorm:"column:author" validate:"required"`
	Summary  string            `json:"summary" gorm:"column:summary"`
	Content  string            `json:"content" gorm:"column:content" validate:"required"`
	CreateBy string            `json:"create_by" gorm:"column:create_by"`
	Tags     map[string]string `json:"tags" gorm:"column:tags;serializer:json;serializer:json"`
}

func (c *CreateBlogRequest) Validate() error {
	return common.Validate(c)
}

func NewChangeBlogStatusRequest() *ChangeBlogStatusRequest {
	return &ChangeBlogStatusRequest{
		Status: STATUS_DRAFT,
	}
}

type ChangeBlogStatusRequest struct {
	PublishedAt int64  `json:"published_at" gorm:"column:published_at"`
	Status      Status `json:"status" gorm:"status"`
}

func (req *ChangeBlogStatusRequest) SetStatus(s Status) {
	req.Status = s
	switch req.Status {
	case STATUS_PUBLISH:
		req.PublishedAt = time.Now().Unix()
	}
}

func NewBlog() *Blog {
	return &Blog{
		Mate: common.NewMate(),
		CreateBlogRequest: &CreateBlogRequest{
			Tags: map[string]string{},
		},
		ChangeBlogStatusRequest: &ChangeBlogStatusRequest{
			Status: STATUS_DRAFT,
		},
	}
}

type Blog struct {
	*common.Mate
	*CreateBlogRequest
	*ChangeBlogStatusRequest
}

func NewBlogSet() *BlogSet {
	return &BlogSet{
		Items: []*Blog{},
	}
}

type BlogSet struct {
	Total int64   `json:"total"`
	Items []*Blog `json:"items"`
}
