package dtos

import "github.com/vitorpereiraa/go-microservices-starter/TodoList/domain"

func ItemToDto(item domain.Item) ItemDTO {
	return ItemDTO{
		Description: item.Description(),
		DueDate:     item.DueDate(),
	}
}

func ItemListToDtoList(items []domain.Item) (itemsDto []ItemDTO) {
	for _, item := range items {
		itemsDto = append(itemsDto, ItemToDto(item))
	}
	return
}

func ItemDtoToItem(itemDto ItemDTO) (Item domain.Item) {
	return domain.NewItem(itemDto.Description, itemDto.DueDate)
}

func ItemDtoListToItemList(itemsDto []ItemDTO) (items []domain.Item) {
	for _, item := range itemsDto {
		items = append(items, ItemDtoToItem(item))
	}
	return
}
