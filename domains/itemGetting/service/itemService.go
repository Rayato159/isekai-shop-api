package service

import (
	_itemGettingModel "github.com/Rayato159/isekai-shop-api/domains/itemGetting/model"
)

type ItemService interface {
	ItemListing(itemFilter *_itemGettingModel.ItemFilter) (*_itemGettingModel.ItemResult, error)
}
