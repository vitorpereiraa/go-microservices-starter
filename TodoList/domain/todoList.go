package domain

import "gorm.io/gorm"

type TodoList struct {
	gorm.Model
	name string
	description string
	items       []Item
}

func NewTodoList(name string, descrition string, items []Item) TodoList {
	return TodoList{
		name : name,
		description: descrition,
		items:       items,
	}
}

func (t *TodoList) Name() string {
	return t.name
}

func (t *TodoList) SetName(name string) {
	t.name = name
}

func (t *TodoList) Description() string {
	return t.description
}

func (t *TodoList) SetDescription(description string) {
	t.description = description
}

func (t *TodoList) Items() []Item {
	return t.items
}

func (t *TodoList) SetItems(items []Item) {
	t.items = items
}
