package impl_test

import (
	"blog/apps/token"
	"blog/ioc"
	"blog/test"
	"context"
	"testing"
)

var (
	ctx          = context.Background()
	TokenService token.Service
)

func init() {
	test.DevelopmentSetup()
	TokenService = ioc.Contorller.Get(token.AppName).(token.Service)
}

func TestIssue(t *testing.T) {
	req := token.NewIssueTokenRequest("ne", "ne")
	tk, err := TokenService.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)

}

func TestRevoke(t *testing.T) {
	req := token.NewRevokeTokenRequest("d8d411gn69ahgs70kt2g", "d8d411gn69ahgs70kt30")
	tk, err := TokenService.RevokeToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func TestValidate(t *testing.T) {
	req := token.NewValidateTokenRequest("d8d42eon69aj1r504t70")
	tk, err := TokenService.ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}
