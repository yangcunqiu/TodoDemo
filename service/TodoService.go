package service

import (
	"TodoDemo/common"
	"TodoDemo/convertor"
	"TodoDemo/dao"
	"TodoDemo/model/dto"
	"TodoDemo/model/enums"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func AddTodo(todoDTO dto.TodoDTO, c *gin.Context) {
	todoDTO.Status = enums.UNFINISHED
	todo := convertor.DTOToEntity(todoDTO)
	// 保存
	dao.AddTodo(todo)
}

func GetInfoById(idStr string, c *gin.Context) dto.TodoDTO {
	// 校验
	if idStr == "" {
		common.FailureNotError(c, "id不能为空")
	}
	id, err := strconv.Atoi(idStr)
	common.Failure(err, c, "id解析失败, 请重试")
	// 查询
	todo, result := dao.GetTodoById(id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		common.FailureNotError(c, "todo未找到")
	}
	return convertor.EntityToDTO(todo)
}

func ListTodo(c *gin.Context) []dto.TodoDTO {
	return convertor.ListEntityToDTO(dao.ListTodo())
}

func UpdateTodo(todoDTO dto.TodoDTO, c *gin.Context) {
	// 校验
	GetInfoById(strconv.Itoa(todoDTO.ID), c)
	// 更新
	dao.UpdateTodo(convertor.DTOToEntity(todoDTO))
}

func DeletedTodo(idStr string, c *gin.Context) {
	// 校验
	GetInfoById(idStr, c)
	id, err := strconv.Atoi(idStr)
	common.Failure(err, c, "id解析失败, 请重试")
	// 删除
	dao.DeletedTodo(id)
}

func FinishTodo(idStr string, c *gin.Context) {
	// 校验
	GetInfoById(idStr, c)
	id, err := strconv.Atoi(idStr)
	common.Failure(err, c, "id解析失败, 请重试")
	// 更新todo状态为已完成
	finishTodoDTO := dto.TodoDTO{
		ID:     id,
		Status: enums.FINISHED,
	}
	UpdateTodo(finishTodoDTO, c)

}

func ResetTodo(idStr string, c *gin.Context) {
	// 校验
	GetInfoById(idStr, c)
	id, err := strconv.Atoi(idStr)
	common.Failure(err, c, "id解析失败, 请重试")
	// 更新todo状态为未完成
	resetTodoDTO := dto.TodoDTO{
		ID:     id,
		Status: enums.UNFINISHED,
	}
	UpdateTodo(resetTodoDTO, c)
}
