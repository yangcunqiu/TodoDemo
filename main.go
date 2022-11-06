package main

import (
	"TodoDemo/controller"
	"TodoDemo/dao"
)

func main() {
	// 初始化数据库
	dao.InitDB()

	// 初始化gin
	controller.InitGin()

	// 初始化路由
	controller.InitRouter()
}
