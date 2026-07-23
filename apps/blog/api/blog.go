package api

import (
	"blog/apps/blog"
	"blog/apps/token"
	"blog/common"
	"blog/exception"
	"blog/middleware"
	"blog/response"

	"github.com/gin-gonic/gin"
)

func (h *BlogApiHandler) Registry(appRouter gin.IRouter) {
	appRouter.GET("/", h.QueryBlog)
	appRouter.GET("/:id", h.DescribeBlog)

	appRouter.Use(middleware.Auth)
	appRouter.POST("/", h.CreateBlog)
	appRouter.PUT("/:id", h.UpdateBlogPut)
	appRouter.PATCH("/:id", h.UpdateBlogPath)
	appRouter.POST("/:id/status", h.UpdateBlogStatus)
	appRouter.DELETE("/:id", h.DeleteBlog)

}

func (h *BlogApiHandler) CreateBlog(ctx *gin.Context) {
	req := blog.NewCreateBlogRequest()
	if err := ctx.Bind(req); err != nil {
		response.Fatal(exception.ErrValidateFailed("%s ErrValidate", err.Error()), ctx)
		return
	}
	if v, ok := ctx.Get(token.GIN_TOKEN_KEY_NAME); ok {
		req.CreateBy = v.(*token.Token).UserName
	}
	ins, err := h.svc.CreateBlog(ctx.Request.Context(), req)
	if err != nil {
		response.Fatal(err, ctx)
		return
	}
	response.Success(ins, ctx)
}
func (h *BlogApiHandler) DeleteBlog(ctx *gin.Context) {
	req := blog.NewDeleteBlogRequest(ctx.Param("id"))
	ins, err := h.svc.DeleteBlog(ctx.Request.Context(), req)
	if err != nil {
		response.Fatal(err, ctx)
		return
	}
	response.Success(ins, ctx)
}
func (h *BlogApiHandler) UpdateBlogPut(ctx *gin.Context) {
	req := blog.NewUpdateBlogRequest(ctx.Param("id"))
	req.UpdateMode = blog.UPDATE_MODE_PUT
	if err := ctx.Bind(req.CreateBlogRequest); err != nil {
		response.Fatal(exception.ErrValidateFailed("%s ErrValidate", err.Error()), ctx)
		return
	}
	ins, err := h.svc.UpdateBlog(ctx, req)
	if err != nil {
		response.Fatal(err, ctx)
		return
	}
	response.Success(ins, ctx)
}
func (h *BlogApiHandler) UpdateBlogPath(ctx *gin.Context) {
	req := blog.NewUpdateBlogRequest(ctx.Param("id"))
	req.UpdateMode = blog.UPDATE_MODE_PATCH
	if err := ctx.Bind(req.CreateBlogRequest); err != nil {
		response.Fatal(exception.ErrValidateFailed("%s ErrValidate", err.Error()), ctx)
		return
	}
	ins, err := h.svc.UpdateBlog(ctx, req)
	if err != nil {
		response.Fatal(err, ctx)
		return
	}
	response.Success(ins, ctx)

}
func (h *BlogApiHandler) UpdateBlogStatus(ctx *gin.Context) {
	req := blog.NewUpdateStatusBlogRequest(ctx.Param("id"))
	if err := ctx.Bind(req.ChangeBlogStatusRequest); err != nil {
		response.Fatal(exception.ErrValidateFailed("%s ErrValidate", err.Error()), ctx)
		return
	}
	ins, err := h.svc.UpdateStatusBlog(ctx.Request.Context(), req)
	if err != nil {
		response.Fatal(err, ctx)
		return
	}
	response.Success(ins, ctx)
}
func (h *BlogApiHandler) QueryBlog(ctx *gin.Context) {
	req := blog.NewQueryBlogRequest()
	req.PageRequest = common.NewPageRequestFromGinCtx(ctx)
	req.KeyWords = ctx.Query("keywords")

	set, err := h.svc.QueryBlog(ctx.Request.Context(), req)
	if err != nil {
		response.Fatal(err, ctx)
		return
	}
	response.Success(set, ctx)
}
func (h *BlogApiHandler) DescribeBlog(ctx *gin.Context) {
	req := blog.NewDesibleBlogRequest(ctx.Param("id"))
	ins, err := h.svc.DesibleBlog(ctx.Request.Context(), req)
	if err != nil {
		response.Fatal(err, ctx)
		return
	}
	response.Success(ins, ctx)
}
