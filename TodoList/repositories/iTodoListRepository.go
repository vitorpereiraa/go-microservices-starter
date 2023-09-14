package repositories

import "github.com/vitorpereiraa/go-microservices-starter/TodoList/domain"

type ITodoListRepository interface {
	FindTodoLists() []domain.TodoList
	FindTodoListByName(name string) (domain.TodoList, error)
	CreateTodoList(todoList domain.TodoList) (domain.TodoList, error)
	DeleteTodoListByName(name string) (domain.TodoList, error)
}
