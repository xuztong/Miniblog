package middleware

import (
	"blog/apps/token"
	"blog/ioc"
	"blog/response"

	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	accessToken, _ := ctx.Cookie(token.COOKIE_TOKEN_KEY)
	tk, err := ioc.Contorller.Get(token.AppName).(token.Service).ValidateToken(ctx.Request.Context(), token.NewValidateTokenRequest(accessToken))
	if err != nil {
		response.Fatal(token.ErrAuthFailed.WithMessage(err.Error()), ctx)
		ctx.Abort()
	} else {
		ctx.Set(token.GIN_TOKEN_KEY_NAME, tk)
		ctx.Next()
	}
}
