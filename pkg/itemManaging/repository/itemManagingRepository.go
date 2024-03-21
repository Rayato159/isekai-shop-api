package repository

import entities "github.com/Rayato159/isekai-shop-api/entities"

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
	Editing(itemID uint64, updateItemDto *entities.ItemEditingDto) (uint64, error)
	Archiving(itemID uint64) error // Soft delete
}
