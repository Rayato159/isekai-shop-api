package service

import (
	_itemModel "github.com/Rayato159/isekai-shop-api/domains/item/model"
)

type ItemService interface {
	ItemListing(itemFilter *_itemModel.ItemFilter) (*_itemModel.ItemResult, error)
}
