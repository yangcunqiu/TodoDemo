package controller

func InitRouter() {

	todoGroup := r.Group("/todo")
	{
		todoGroup.POST("/add", AddTodo)
		todoGroup.GET("/info/:id", GetInfoById)
		todoGroup.GET("/list", ListTodo)
		todoGroup.POST("/update", UpdateTodo)
		todoGroup.POST("/deleted/:id", DeletedTodo)
		todoGroup.POST("/finish/:id", FinishTodo)
		todoGroup.POST("/reset/:id", ResetTodo)
	}

	// 启动
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
