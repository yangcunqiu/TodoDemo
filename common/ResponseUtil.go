package common

import (
	"TodoDemo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Failure(err error, c *gin.Context, message string) {
	if err != nil {
		resp := model.Response[interface{}]{
			Code:    FAILURE,
			Message: message,
		}
		c.JSON(http.StatusOK, resp)
	}
	c.Next()
}

func FailureNotError(c *gin.Context, message string) {
	resp := model.Response[interface{}]{
		Code:    FAILURE,
		Message: message,
	}
	c.JSON(http.StatusOK, resp)
	c.Abort()
}

func Success(c *gin.Context, data ...interface{}) {
	resp := model.Response[interface{}]{
		Code:    SUCCESS,
		Message: "成功",
		Data:    data,
	}
	c.JSON(http.StatusOK, resp)
	c.Next()
}
