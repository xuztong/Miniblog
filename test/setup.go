package test

import (
	_ "blog/apps"
	"blog/conf"
	"blog/ioc"
)

func DevelopmentSetup() {
	conf.C()
	

	ioc.Contorller.Init()
	ioc.Api.Init()
}
