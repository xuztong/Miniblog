package user

import (
	"blog/common"
	"context"
)

const(
	AppName="user"
)

type Service interface {
	CreateUser(context.Context, *CreateUserReuqest) (*User, error)
	QueryUser(context.Context, *QueryUserRequest) (*UserSet, error)
}

func NewQueryUserRequest() *QueryUserRequest {
	return &QueryUserRequest{
		PageRequest: common.NewPageRequest(),
	}
}

type QueryUserRequest struct {
	Username string
	*common.PageRequest
}
