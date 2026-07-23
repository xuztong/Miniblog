package api

import (
	"blog/apps/token"
	"blog/conf"
	"blog/ioc"
	"blog/response"

	"github.com/gin-gonic/gin"
)

func init() {
	ioc.Api.Registry(token.AppName, &TokenApiHander{})
}

type TokenApiHander struct {
	token token.Service
}

func (h *TokenApiHander) Init() error {
	h.token = ioc.Contorller.Get(token.AppName).(token.Service)
	TokenRouter := conf.C().App.GinRootServer().Group("tokens")
	h.Registyr(TokenRouter)
	return nil
}

func (h *TokenApiHander) Registyr(AppRoter gin.IRouter) {
	AppRoter.POST("/", h.Login)
	AppRoter.DELETE("/",h.Logout)
}

func (h *TokenApiHander) Login(c *gin.Context) {
	//获取name,password
	req := token.NewIssueTokenRequest("", "")
	if err := c.BindJSON(req); err != nil {
		response.Fatal(err, c)
		return
	}
	//颁发token
	tk, err := h.token.IssueToken(c.Request.Context(), req)
	if err != nil {
		response.Fatal(err, c)
		return
	}
	c.SetCookie(
		token.COOKIE_TOKEN_KEY,
		tk.AccessToken,
		tk.AccessTokenExpiredAt,
		"/",
		conf.C().App.Domain,
		false,
		true,
	)
	response.Success(tk, c)
}

func (h *TokenApiHander) Logout(c *gin.Context) {
	//获取 access refresh
	at, err := c.Cookie(token.COOKIE_TOKEN_KEY)
	if err != nil {
		response.Fatal(err, c)
		return
	}
	rt := c.GetHeader(token.REFRESH_HEADER_KEY)

	req := token.NewRevokeTokenRequest(at, rt)
	tk, err := h.token.RevokeToken(c.Request.Context(), req)
	if err != nil {
		response.Fatal(err, c)
		return
	}
	response.Success(tk, c)

}
