package services

import (
	"github.com/vitorpereiraa/go-microservices-starter/TodoList/dtos"
	"github.com/vitorpereiraa/go-microservices-starter/TodoList/repositories"
)

type TodoListService struct {
	repo repositories.ITodoListRepository
}

func NewTodoListService(todoListRepo repositories.ITodoListRepository) *TodoListService {
	return &TodoListService{
		repo: todoListRepo,
	}
}

func (s TodoListService) FindTodoLists() []dtos.TodoListDTO {
	todoLists := s.repo.FindTodoLists()
	return dtos.TodoListListToDtoList(todoLists)
}

func (s TodoListService) FindTodoListByName(id string) (dtos.TodoListDTO, error) {
	todoList, err := s.repo.FindTodoListByName(id)
	return dtos.TodoListToDto(todoList), err
}

func (s TodoListService) CreateTodoList(todoListDTO dtos.TodoListDTO) (dtos.TodoListDTO, error) {
	todoList := dtos.TodoListDtoToTodoList(todoListDTO)
	createdTodoList, err := s.repo.CreateTodoList(todoList)
	return dtos.TodoListToDto(createdTodoList), err
}

func (s TodoListService) DeleteTodoListByName(id string) (dtos.TodoListDTO, error) {
	todoList, err := s.repo.DeleteTodoListByName(id)
	return dtos.TodoListToDto(todoList), err
}
