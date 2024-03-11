package service

import (
	_itemEntity "github.com/Rayato159/isekai-shop-api/modules/item/entity"
	_itemModel "github.com/Rayato159/isekai-shop-api/modules/item/model"
	_itemRepository "github.com/Rayato159/isekai-shop-api/modules/item/repository"
)

type itemServiceImpl struct {
	itemRepository _itemRepository.ItemRepository
}

func NewItemServiceImpl(itemRepository _itemRepository.ItemRepository) ItemService {
	return &itemServiceImpl{
		itemRepository: itemRepository,
	}
}

func (s *itemServiceImpl) ItemListing(itemFilter *_itemModel.ItemFilter) (*_itemModel.ItemResult, error) {
	itemFilterDto := &_itemEntity.ItemFilterDto{
		Name:        itemFilter.Name,
		Description: itemFilter.Description,
		PaginateDto: _itemEntity.PaginateDto{
			Page: itemFilter.Page,
			Size: itemFilter.Size,
		},
	}

	itemEntityList, err := s.itemRepository.ItemListing(itemFilterDto)
	if err != nil {
		return nil, err
	}

	totalItems, err := s.itemRepository.ItemCounting(itemFilterDto)
	if err != nil {
		return nil, err
	}

	size := itemFilter.Paginate.Size
	page := itemFilter.Paginate.Page
	totalPage := s.totalPageCalculation(totalItems, size)

	result := s.buildItemResultsResponse(itemEntityList, page, totalPage)

	return result, nil
}

func (s *itemServiceImpl) totalPageCalculation(totalItems, size int64) int64 {
	totalPage := totalItems / size

	if totalItems%size != 0 {
		totalPage++
	}

	return totalPage
}

func (s *itemServiceImpl) buildItemResultsResponse(itemEntityList []*_itemEntity.Item, page, totalPage int64) *_itemModel.ItemResult {
	items := make([]*_itemModel.Item, 0)

	for _, itemEntity := range itemEntityList {
		items = append(items, itemEntity.ToItemModel())
	}

	return &_itemModel.ItemResult{
		Items: items,
		Paginate: _itemModel.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}
