package service

import (
	_itemEntity "github.com/Rayato159/isekai-shop-api/modules/item/entity"
	_itemModel "github.com/Rayato159/isekai-shop-api/modules/item/model"
	_itemRepository "github.com/Rayato159/isekai-shop-api/modules/item/repository"
)

type adminServiceImpl struct {
	itemRepository _itemRepository.ItemRepository
}

func NewAdminServiceImpl(itemRepository _itemRepository.ItemRepository) AdminService {
	return &adminServiceImpl{
		itemRepository: itemRepository,
	}
}

func (s *adminServiceImpl) ItemCreating(itemCreatingReq *_itemModel.ItemCreatingReq) (*_itemModel.Item, error) {
	item := &_itemEntity.Item{
		AdminID:     &itemCreatingReq.AdminID,
		Name:        itemCreatingReq.Name,
		Description: itemCreatingReq.Description,
		Picture:     itemCreatingReq.Picture,
		Price:       itemCreatingReq.Price,
	}

	itemEntity, err := s.itemRepository.ItemCreating(item)
	if err != nil {
		return nil, err
	}

	return itemEntity.ToItemModel(), nil
}

func (s *adminServiceImpl) ItemEditing(itemID uint64, updateItemReq *_itemModel.ItemEditingReq) (*_itemModel.Item, error) {
	updateItemDto := &_itemEntity.ItemEditingDto{
		AdminID:     &updateItemReq.AdminID,
		Name:        updateItemReq.Name,
		Description: updateItemReq.Description,
		Picture:     updateItemReq.Picture,
		Price:       updateItemReq.Price,
	}

	updatedItemID, err := s.itemRepository.ItemEditing(itemID, updateItemDto)
	if err != nil {
		return nil, err
	}

	itemEntity, err := s.itemRepository.FindItemByID(updatedItemID)
	if err != nil {
		return nil, err
	}

	return itemEntity.ToItemModel(), nil
}

func (s *adminServiceImpl) ItemArchiving(itemID uint64) error {
	return s.itemRepository.ItemArchiving(itemID)
}
