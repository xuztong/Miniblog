package impl_test

import (
	"blog/apps/blog"
	"blog/ioc"
	"blog/test"
	"context"
)

var (
	serviceImpl blog.Service
	ctx         = context.Background()
)

func init() {
	test.DevelopmentSetup()
	serviceImpl = ioc.Contorller.Get(blog.AppName).(blog.Service)
}
