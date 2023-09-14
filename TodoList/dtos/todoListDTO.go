package dtos

type TodoListDTO struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Items       []ItemDTO `json:"items"`
}
