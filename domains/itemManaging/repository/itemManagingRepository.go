package repository

import entities "github.com/Rayato159/isekai-shop-api/entities"

type ItemManagingRepository interface {
	ItemCreating(itemEntity *entities.Item) (*entities.Item, error)
	ItemEditing(itemID uint64, updateItemDto *entities.ItemEditingDto) (uint64, error)
	ItemArchiving(itemID uint64) error // Soft delete
}
