package impl_test

import (
	"blog/apps/user"
	"blog/ioc"
	"blog/test"
	"context"
	"testing"
)

var (
	ctx         = context.Background()
	UserService user.Service
)

func init() {
	test.DevelopmentSetup()
	UserService = ioc.Contorller.Get(user.AppName).(user.Service)
}

func TestCreate(t *testing.T) {
	ne := user.NewCreateUserReuqest()
	ne.Username = "root"
	ne.Password = "root"
	us, err := UserService.CreateUser(ctx, ne)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(us)
}

func TestQuery(t *testing.T) {
	ne := user.NewQueryUserRequest()
	ne.Username = "ne"
	set, err := UserService.QueryUser(ctx, ne)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set.Items[0].Password)
}
