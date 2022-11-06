package dao

import (
	"TodoDemo/model/entity"
	"gorm.io/gorm"
)

func AddTodo(todo entity.Todo) {
	db.Create(&todo)
}

func GetTodoById(id int) (entity.Todo, *gorm.DB) {
	var todo entity.Todo
	result := db.First(&todo, id)
	return todo, result
}

func ListTodo() []entity.Todo {
	todoList := make([]entity.Todo, 0)
	db.Find(&todoList)
	return todoList
}

func UpdateTodo(todo entity.Todo) {
	db.Debug().Model(&todo).Select("status").Updates(todo)
}

func DeletedTodo(id int) {
	db.Delete(&entity.Todo{}, id)
}
