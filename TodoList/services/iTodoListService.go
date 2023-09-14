package services

import "github.com/vitorpereiraa/go-microservices-starter/TodoList/dtos"

type ITodoListService interface {
	FindTodoLists() []dtos.TodoListDTO
	FindTodoListByName(name string) (dtos.TodoListDTO, error)
	CreateTodoList(todoListDTO dtos.TodoListDTO) (dtos.TodoListDTO, error)
	DeleteTodoListByName(name string) (dtos.TodoListDTO, error)
}