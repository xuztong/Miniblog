package token

import "context"

const(
	AppName="token"
)

type Service interface {
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	RevokeToken(context.Context, *RevokeTokenRequest) (*Token, error)
	ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error)
}

func NewIssueTokenRequest(username string, password string) *IssueTokenRequest {
	return &IssueTokenRequest{
		Username: username,
		Password: password,
	}
}

type IssueTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsMember bool   `json:"is_member"`
}

func NewRevokeTokenRequest(at string, rt string) *RevokeTokenRequest {
	return &RevokeTokenRequest{
		AccessToken:  at,
		RefreshToken: rt,
	}
}

type RevokeTokenRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewValidateTokenRequest(at string) *ValidateTokenRequest {
	return &ValidateTokenRequest{
		AccessToken: at,
	}
}

type ValidateTokenRequest struct {
	AccessToken string `json:"access_token"`
}
