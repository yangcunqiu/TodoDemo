package dto

type TodoDTO struct {
	ID int `json:"id"`
	// 标题
	Title string `json:"title" binding:"required"`
	// 状态 0 未完成 1 已完成
	Status byte `json:"status"`
}
