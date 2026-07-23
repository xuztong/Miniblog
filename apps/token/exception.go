package token

import (
	"blog/exception"
	"net/http"
)

var (
	ErrAuthFailed          = exception.NewApiException(50001, "用户名或者密码不正确").WithHttpCode(http.StatusUnauthorized)
	ErrAccessTokenExpired  = exception.NewApiException(50002, "AccessToken过期")
	ErrRefreshTokenExpired = exception.NewApiException(50003, "RefreshToken过期")
	ErrPermissionDeny      = exception.NewApiException(50004, "当前角色无权限访问该接口").WithHttpCode(http.StatusForbidden)
)
