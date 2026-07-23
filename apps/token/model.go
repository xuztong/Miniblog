package token

import (
	"blog/apps/user"
	"time"

	"github.com/rs/xid"
)

func NewToken(u *user.User) *Token {
	return &Token{
		UserId:                u.Id,
		UserName:              u.Username,
		AccessToken:           xid.New().String(),
		AccessTokenExpiredAt:  3600,
		RefreshToken:          xid.New().String(),
		RefreshTokenExpiredAt: 3600 * 4,
		CreatedAt:             time.Now().Unix(),
		Role:                  u.Role,
	}
}

func TokenDefault() *Token {
	return &Token{}
}

type Token struct {
	UserId                int       `json:"user_id" gorm:"column:user_id"`
	UserName              string    `json:"username" gorm:"column:username"`
	AccessToken           string    `json:"access_token" gorm:"column:access_token"`
	AccessTokenExpiredAt  int       `json:"access_token_expired_at" gorm:"column:access_token_expired_at"`
	RefreshToken          string    `json:"refresh_token" gorm:"column:refresh_token"`
	RefreshTokenExpiredAt int       `json:"refresh_token_expired_at" gorm:"column:refresh_token_expired_at"`
	CreatedAt             int64     `json:"created_at" gorm:"column:created_at"`
	UpdatedAt             int64     `json:"updated_at" gorm:"column:updated_at"`
	Role                  user.Role `json:"role" gorm:"-"`
}

func (t *Token) IssueTime() time.Time {
	return time.Unix(t.CreatedAt, 0)
}

func (t *Token) AccessTokenDuration() time.Duration {
	return time.Duration(t.AccessTokenExpiredAt) * time.Second
}

func (t *Token) RefreshTokenDuration() time.Duration {
	return time.Duration(t.RefreshTokenExpiredAt) * time.Second
}

func (t *Token) AccessTokenIsExpiredAt() error {
	expiredTime := t.IssueTime().Add(t.AccessTokenDuration())
	if time.Now().After(expiredTime) {
		return ErrAccessTokenExpired
	}
	return nil
}

func (t *Token) RefreshTokenIsExpiredAt() error {
	expiredTime := t.IssueTime().Add(t.RefreshTokenDuration())
	if time.Now().After(expiredTime) {
		return ErrRefreshTokenExpired
	}
	return nil
}
