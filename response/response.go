package response

import (
	"blog/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(data any, c *gin.Context) {
	c.JSON(http.StatusOK, data)
}

func Fatal(err error, c *gin.Context) {
	httpcode := http.StatusInternalServerError
	if v, ok := err.(*exception.ApiException); ok {
		if v.HttpCode != 0 {
			httpcode = v.HttpCode
		}
	} else {
		exception.ErrServerInternal("ErrServerInternal: %s",err.Error())
	}
	c.JSON(httpcode, err)
	c.Abort()
}
