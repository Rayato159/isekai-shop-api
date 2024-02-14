package repository

import (
	_itemEntity "github.com/Rayato159/isekai-shop-api/modules/item/entity"
)

type ItemRepository interface {
	FindItems(itemFilterDto *_itemEntity.ItemFilterDto) ([]*_itemEntity.Item, error)
	CountItems(itemFilterDto *_itemEntity.ItemFilterDto) (int64, error)
}
