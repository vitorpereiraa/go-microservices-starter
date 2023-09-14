package dtos

import "github.com/vitorpereiraa/go-microservices-starter/TodoList/domain"

func TodoListToDto(todoList domain.TodoList) TodoListDTO {
	return TodoListDTO{
		Name: todoList.Name(),
		Description: todoList.Description(),
		Items:       ItemListToDtoList(todoList.Items()),
	}
}

func TodoListListToDtoList(todoListList []domain.TodoList) (todoListDtos []TodoListDTO) {
	for _, todoList := range todoListList {
		todoListDtos = append(todoListDtos, TodoListToDto(todoList))
	}
	return
}

func TodoListDtoToTodoList(todoListDTO TodoListDTO) domain.TodoList {
	return domain.NewTodoList(todoListDTO.Name,todoListDTO.Description, ItemDtoListToItemList(todoListDTO.Items))
}
