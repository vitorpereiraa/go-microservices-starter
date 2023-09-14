package repositories

import (
	"github.com/vitorpereiraa/go-microservices-starter/TodoList/domain"
	"gorm.io/gorm"
)

type TodoListRepositoryPostgres struct {
	db *gorm.DB 
}

func NewTodoListRepositoryPostgres(db *gorm.DB) *TodoListRepositoryPostgres {
	repo  := &TodoListRepositoryPostgres{
		db: db,
	}

	bootstrapPostgres(repo)

	return repo
}

func (t *TodoListRepositoryPostgres) FindTodoLists() (list []domain.TodoList) {
	t.db.Find(&list)
	return
}

func (t *TodoListRepositoryPostgres) FindTodoListByName(name string) (todoList domain.TodoList, err error) {
	t.db.Where(map[string]interface{}{"name": name}).First(&todoList)
	return 
}

func (t *TodoListRepositoryPostgres) CreateTodoList(todoList domain.TodoList) (domain.TodoList, error) {
	t.db.Create(&todoList)
	return todoList, nil 
}

func (t *TodoListRepositoryPostgres) DeleteTodoListByName(name string) (todoList domain.TodoList, err error) {
	t.db.Where(map[string]interface{}{"name": name}).Delete(&todoList)
	return todoList, nil
}

func bootstrapPostgres(todoListRepository *TodoListRepositoryPostgres) {
	var items []domain.Item

	items = append(items, domain.NewItem("Do Homework", "Until Tomorrow"))

	var todoList domain.TodoList = domain.NewTodoList("School Todo List", "School Tasks", items)

	todoListRepository.db.Create(&todoList)
}
