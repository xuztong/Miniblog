package impl

import (
	"blog/apps/token"
	"blog/apps/user"
	"blog/conf"
	"blog/ioc"
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func init() {
	ioc.Contorller.Registry(token.AppName, &TokenServiceImpl{})
}

type TokenServiceImpl struct {
	db    *gorm.DB
	cache *redis.Client
	user  user.Service
}

func (i *TokenServiceImpl) Init() error {
	i.db = conf.C().Mysql.GetDB()
	i.cache = conf.C().Redis.GetClient()
	i.user = ioc.Contorller.Get(user.AppName).(user.Service)
	return nil
}

func tokenCacheKey(accessToken string) string {
	return "token:access" + accessToken
}

func (i *TokenServiceImpl) IssueToken(ctx context.Context, in *token.IssueTokenRequest) (*token.Token, error) {
	//查询用户
	req := user.NewQueryUserRequest()
	req.Username = in.Username
	us, err := i.user.QueryUser(ctx, req)
	if err != nil {
		return nil, err
	}
	//对比密码
	use := us.Items[0]
	if err := use.CheckPassword(in.Password); err != nil {
		return nil, err
	}
	//颁发令牌
	tk := token.NewToken(use)

	data, err := json.Marshal(tk)
	if err != nil {
		return nil, err
	}
	ttl := time.Duration(tk.AccessTokenExpiredAt) * time.Second

	key := tokenCacheKey(tk.AccessToken)
	if err := i.cache.Set(ctx, key, data, ttl).Err(); err != nil {
		return nil, err
	}
	return tk, err
}
func (i *TokenServiceImpl) RevokeToken(ctx context.Context, in *token.RevokeTokenRequest) (*token.Token, error) {
	key := tokenCacheKey(in.AccessToken)
	data, err := i.cache.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, token.ErrAuthFailed
	}
	if err != nil {
		return nil, err
	}
	tk := token.TokenDefault()
	if err := json.Unmarshal(data, tk); err != nil {
		return nil, err
	}
	if tk.RefreshToken != in.RefreshToken {
		return nil, token.ErrAuthFailed
	}
	if err := i.cache.Del(ctx, key).Err(); err != nil {
		return nil, err
	}
	return tk, nil
}
func (i *TokenServiceImpl) ValidateToken(ctx context.Context, in *token.ValidateTokenRequest) (*token.Token, error) {

	key := tokenCacheKey(in.AccessToken)

	data, err := i.cache.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, token.ErrAccessTokenExpired
	}

	if err != nil {
		return nil, err
	}
	tk := token.TokenDefault()
	if err := json.Unmarshal(data, tk); err != nil {
		return nil, err
	}

	return tk, nil

}
