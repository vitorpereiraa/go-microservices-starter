package domain

type Item struct {
	description string
	dueDate     string
}

func NewItem(description string, dueDate string) Item {
	return Item{
		description: description,
		dueDate:     dueDate,
	}
}

func (i *Item) Description() string {
	return i.description
}

func (i *Item) SetDescription(description string) {
	i.description = description
}

func (i *Item) DueDate() string {
	return i.dueDate
}

func (i *Item) SetDueDate(dueDate string) {
	i.dueDate = dueDate
}
