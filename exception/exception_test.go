package exception_test

import (
	"blog/exception"
	"testing"
)

func E() error {
	return exception.NewApiException(50000, "hello")
}

func TestErr(t *testing.T) {
	err := E()
	t.Log(err)

	if v, ok := err.(*exception.ApiException); ok {
		t.Log(v.Code)
		t.Log(v.Message)
	}
}
