package service

import (
	_itemShopModel "github.com/Rayato159/isekai-shop-api/domains/itemShop/model"
	_itemShopRepository "github.com/Rayato159/isekai-shop-api/domains/itemShop/repository"
	entities "github.com/Rayato159/isekai-shop-api/entities"
)

type itemServiceImpl struct {
	itemRepository _itemShopRepository.ItemShopRepository
}

func NewItemServiceImpl(itemRepository _itemShopRepository.ItemShopRepository) ItemService {
	return &itemServiceImpl{
		itemRepository: itemRepository,
	}
}

func (s *itemServiceImpl) Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error) {
	itemFilterDto := &entities.ItemFilterDto{
		Name:        itemFilter.Name,
		Description: itemFilter.Description,
		PaginateDto: entities.PaginateDto{
			Page: itemFilter.Page,
			Size: itemFilter.Size,
		},
	}

	itemEntityList, err := s.itemRepository.Listing(itemFilterDto)
	if err != nil {
		return nil, err
	}

	totalItems, err := s.itemRepository.Counting(itemFilterDto)
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

func (s *itemServiceImpl) buildItemResultsResponse(itemEntityList []*entities.Item, page, totalPage int64) *_itemShopModel.ItemResult {
	items := make([]*_itemShopModel.Item, 0)

	for _, itemEntity := range itemEntityList {
		items = append(items, itemEntity.ToItemModel())
	}

	return &_itemShopModel.ItemResult{
		Items: items,
		Paginate: _itemShopModel.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}
