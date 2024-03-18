package service

import (
	_itemGettingModel "github.com/Rayato159/isekai-shop-api/domains/itemGetting/model"
	_itemGettingRepository "github.com/Rayato159/isekai-shop-api/domains/itemGetting/repository"
	entities "github.com/Rayato159/isekai-shop-api/entities"
)

type itemServiceImpl struct {
	itemRepository _itemGettingRepository.ItemGettingRepository
}

func NewItemServiceImpl(itemRepository _itemGettingRepository.ItemGettingRepository) ItemService {
	return &itemServiceImpl{
		itemRepository: itemRepository,
	}
}

func (s *itemServiceImpl) Listing(itemFilter *_itemGettingModel.ItemFilter) (*_itemGettingModel.ItemResult, error) {
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

func (s *itemServiceImpl) buildItemResultsResponse(itemEntityList []*entities.Item, page, totalPage int64) *_itemGettingModel.ItemResult {
	items := make([]*_itemGettingModel.Item, 0)

	for _, itemEntity := range itemEntityList {
		items = append(items, itemEntity.ToItemModel())
	}

	return &_itemGettingModel.ItemResult{
		Items: items,
		Paginate: _itemGettingModel.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}
