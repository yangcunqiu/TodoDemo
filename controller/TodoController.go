package controller

import (
	"TodoDemo/common"
	"TodoDemo/model/dto"
	"TodoDemo/service"
	"github.com/gin-gonic/gin"
)

// AddTodo 添加
func AddTodo(c *gin.Context) {
	// 绑定
	var todoDTO dto.TodoDTO
	err := c.ShouldBindJSON(&todoDTO)
	common.Failure(err, c, common.PARAMSBINDERROR)
	// 保存
	service.AddTodo(todoDTO, c)
	common.Success(c)
}

// GetInfoById 通过id获取todo
func GetInfoById(c *gin.Context) {
	// 绑定
	id := c.Param("id")
	todoDTO := service.GetInfoById(id, c)
	common.Success(c, todoDTO)
}

func ListTodo(c *gin.Context) {
	todoDTOList := service.ListTodo(c)
	common.Success(c, todoDTOList)
}

// UpdateTodo 更新
func UpdateTodo(c *gin.Context) {
	// 绑定
	var todoDTO dto.TodoDTO
	err := c.ShouldBindJSON(&todoDTO)
	common.Failure(err, c, common.PARAMSBINDERROR)
	// 更新
	service.UpdateTodo(todoDTO, c)
	common.Success(c)
}

func DeletedTodo(c *gin.Context) {
	// 绑定
	id := c.Param("id")
	service.DeletedTodo(id, c)
	common.Success(c)
}

func FinishTodo(c *gin.Context) {
	// 绑定
	id := c.Param("id")
	// 完成todo
	service.FinishTodo(id, c)
	common.Success(c)
}

func ResetTodo(c *gin.Context) {
	// 绑定
	id := c.Param("id")
	// 重置todo
	service.ResetTodo(id, c)
	common.Success(c)
}
