package api

import (
	"blog/apps/blog"
	"blog/conf"
	"blog/ioc"
)

func init() {
	ioc.Api.Registry(blog.AppName, &BlogApiHandler{})
}

type BlogApiHandler struct {
	svc blog.Service
}

func (h *BlogApiHandler) Init() error {
	h.svc = ioc.Contorller.Get(blog.AppName).(blog.Service)
	subRouter := conf.C().App.GinRootServer().Group("blogs")
	h.Registry(subRouter)
	return nil
}
