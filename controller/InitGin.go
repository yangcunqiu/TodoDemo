package controller

import "github.com/gin-gonic/gin"

var r *gin.Engine

func InitGin() {
	// 获取引擎
	r = gin.Default()
}
