package service

import (
	_itemManagingModel "github.com/Rayato159/isekai-shop-api/domains/itemManaging/model"
	_itemManagingRepository "github.com/Rayato159/isekai-shop-api/domains/itemManaging/repository"
	_itemShopModel "github.com/Rayato159/isekai-shop-api/domains/itemShop/model"
	_itemShopRepository "github.com/Rayato159/isekai-shop-api/domains/itemShop/repository"
	entities "github.com/Rayato159/isekai-shop-api/entities"
)

type itemManagingServiceImpl struct {
	itemManagingRepository _itemManagingRepository.ItemManagingRepository
	itemShopRepository     _itemShopRepository.ItemShopRepository
}

func NewItemManagingServiceImpl(
	itemManagingRepository _itemManagingRepository.ItemManagingRepository,
	itemShopRepository _itemShopRepository.ItemShopRepository,
) ItemManagingService {
	return &itemManagingServiceImpl{
		itemManagingRepository,
		itemShopRepository,
	}
}

func (s *itemManagingServiceImpl) ItemCreating(itemCreatingReq *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error) {
	item := &entities.Item{
		AdminID:     &itemCreatingReq.AdminID,
		Name:        itemCreatingReq.Name,
		Description: itemCreatingReq.Description,
		Picture:     itemCreatingReq.Picture,
		Price:       itemCreatingReq.Price,
	}

	itemEntity, err := s.itemManagingRepository.ItemCreating(item)
	if err != nil {
		return nil, err
	}

	return itemEntity.ToItemModel(), nil
}

func (s *itemManagingServiceImpl) ItemEditing(itemID uint64, updateItemReq *_itemManagingModel.ItemEditingReq) (*_itemShopModel.Item, error) {
	updateItemDto := &entities.ItemEditingDto{
		AdminID:     &updateItemReq.AdminID,
		Name:        updateItemReq.Name,
		Description: updateItemReq.Description,
		Picture:     updateItemReq.Picture,
		Price:       updateItemReq.Price,
	}

	updatedItemID, err := s.itemManagingRepository.ItemEditing(itemID, updateItemDto)
	if err != nil {
		return nil, err
	}

	itemEntity, err := s.itemShopRepository.FindByID(updatedItemID)
	if err != nil {
		return nil, err
	}

	return itemEntity.ToItemModel(), nil
}

func (s *itemManagingServiceImpl) ItemArchiving(itemID uint64) error {
	return s.itemManagingRepository.ItemArchiving(itemID)
}
