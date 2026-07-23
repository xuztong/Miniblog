package api

import (
	"blog/conf"
	"blog/ioc"
	"blog/middleware"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	ioc.Api.Registry("image", &ImageApiHandler{})
}

type ImageApiHandler struct{}

func (h *ImageApiHandler) Init() error {
	// 注册静态文件服务：访问 /uploads/xxx.jpg 直接返回文件
	uploadDir := conf.C().App.UploadDir
	os.MkdirAll(uploadDir, 0755)
	conf.C().App.GinServer().Static("/uploads", uploadDir)

	// 注册上传接口，需要登录才能上传
	r := conf.C().App.GinRootServer().Group("images")
	r.Use(middleware.Auth)
	r.POST("/", h.UploadImage)
	return nil
}

func (h *ImageApiHandler) UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请选择要上传的图片"})
		return
	}

	// 用时间戳做文件名，避免重名
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	uploadDir := conf.C().App.UploadDir
	savePath := filepath.Join(uploadDir, filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "图片保存失败"})
		return
	}

	// 返回可访问的图片 URL
	url := fmt.Sprintf("/uploads/%s", filename)
	c.JSON(http.StatusOK, gin.H{"url": url})
}
