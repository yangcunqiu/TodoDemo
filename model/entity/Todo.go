package entity

import "gorm.io/gorm"

type Todo struct {
	ID int
	// 标题
	Title string
	// 状态 0 未完成 1 已完成
	Status byte
	gorm.Model
}

func (Todo) TableName() string {
	return "todo"
}
