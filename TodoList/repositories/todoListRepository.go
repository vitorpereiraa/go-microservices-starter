package repositories

import (
	"errors"

	"github.com/vitorpereiraa/go-microservices-starter/TodoList/domain"
)

type todoListRepository struct {
	repo []domain.TodoList
}

func NewTodoListRepository() *todoListRepository {
	todoListRepository := &todoListRepository{
		repo: []domain.TodoList{},
	}

	bootstrap(todoListRepository)

	return todoListRepository
}

func (t *todoListRepository) FindTodoLists() []domain.TodoList {
	return t.repo
}

func (t *todoListRepository) FindTodoListByName(name string) (domain.TodoList, error) {
	for i := 0; i < len(t.repo); i++ {
		if t.repo[i].Name() == name {
			return t.repo[i], nil
		}
	}
	return domain.TodoList{}, errors.New("Todo List Not Found")
}

func (t *todoListRepository) CreateTodoList(todoList domain.TodoList) (domain.TodoList, error) {
	t.repo = append(t.repo, todoList)
	return todoList, nil
}

func (t *todoListRepository) DeleteTodoListByName(name string) (domain.TodoList, error) {
	var todoList domain.TodoList

	for i := 0; i < len(t.repo); i++ {
		if t.repo[i].Name() == name {
			t.repo, todoList = remove(t.repo, i)
			return todoList, nil
		}
	}

	return todoList, errors.New("Todo List Not Found")
}

func remove(s []domain.TodoList, i int) ([]domain.TodoList, domain.TodoList) {
	s[i] = s[len(s)-1]
	return s[:len(s)-1], s[i]
}

func bootstrap(todoListRepository *todoListRepository) {
	var items []domain.Item

	items = append(items, domain.NewItem("Do Homework", "Until Tomorrow"))

	var todoList domain.TodoList = domain.NewTodoList("School Todo List", "School Tasks", items)

	todoListRepository.repo = append(todoListRepository.repo, todoList)
}
