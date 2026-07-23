package blog

import (
	"blog/common"
	"context"
)

const (
	AppName = "blogs"
)

type Service interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	DeleteBlog(context.Context, *DeleteBlogRequest) (*Blog, error)
	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)
	DesibleBlog(context.Context, *DesibleBlogRequest) (*Blog, error)
	UpdateStatusBlog(context.Context, *UpdateStatusBlogRequest) (*Blog, error)
}

func NewDeleteBlogRequest(id string) *DeleteBlogRequest {
	return &DeleteBlogRequest{
		BlogId: id,
	}
}

type DeleteBlogRequest struct {
	BlogId string `json:"blog_id"`
}

func NewUpdateBlogRequest(id string) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		BlogId:            id,
		UpdateMode:        UPDATE_MODE_PUT,
		CreateBlogRequest: NewCreateBlogRequest(),
	}
}

type UpdateBlogRequest struct {
	BlogId     string      `json:"blog_id"`
	UpdateMode UPDATE_MODE `json:"update_mode"`
	*CreateBlogRequest
}

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageRequest: common.NewPageRequest(),
	}
}

type QueryBlogRequest struct {
	*common.PageRequest
	KeyWords string `json:"keywords"`
}

func NewDesibleBlogRequest(id string) *DesibleBlogRequest {
	return &DesibleBlogRequest{
		BlogId: id,
	}
}

type DesibleBlogRequest struct {
	BlogId string `json:"blog_id"`
}

func NewUpdateStatusBlogRequest(id string) *UpdateStatusBlogRequest {
	return &UpdateStatusBlogRequest{
		BlogId:                  id,
		ChangeBlogStatusRequest: NewChangeBlogStatusRequest(),
	}
}

type UpdateStatusBlogRequest struct {
	BlogId string `json:"blog_id"`
	*ChangeBlogStatusRequest
}
