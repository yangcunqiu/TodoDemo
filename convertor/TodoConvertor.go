package convertor

import (
	"TodoDemo/model/dto"
	"TodoDemo/model/entity"
)

func DTOToEntity(todoDTO dto.TodoDTO) entity.Todo {
	return entity.Todo{
		ID:     todoDTO.ID,
		Title:  todoDTO.Title,
		Status: todoDTO.Status,
	}
}

func EntityToDTO(todoDTO entity.Todo) dto.TodoDTO {
	return dto.TodoDTO{
		ID:     todoDTO.ID,
		Title:  todoDTO.Title,
		Status: todoDTO.Status,
	}
}

func ListEntityToDTO(todoList []entity.Todo) []dto.TodoDTO {
	todoDTOList := make([]dto.TodoDTO, 0)
	for _, todo := range todoList {
		todoDTOList = append(todoDTOList, EntityToDTO(todo))
	}
	return todoDTOList
}
