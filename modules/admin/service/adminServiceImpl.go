package service

import (
	_itemEntity "github.com/Rayato159/isekai-shop-api/modules/item/entity"
	_itemModel "github.com/Rayato159/isekai-shop-api/modules/item/model"
	_itemRepository "github.com/Rayato159/isekai-shop-api/modules/item/repository"
	"github.com/labstack/echo/v4"
)

type adminServiceImpl struct {
	itemRepository _itemRepository.ItemRepository
	logger         echo.Logger
}

func NewAdminServiceImpl(itemRepository _itemRepository.ItemRepository, logger echo.Logger) AdminService {
	return &adminServiceImpl{
		itemRepository: itemRepository,
		logger:         logger,
	}
}

func (s *adminServiceImpl) CreateItem(createItemReq *_itemModel.CreateItemReq) (*_itemModel.Item, error) {
	item := &_itemEntity.Item{
		AdminID:     &createItemReq.AdminID,
		Name:        createItemReq.Name,
		Description: createItemReq.Description,
		Picture:     createItemReq.Picture,
		Price:       createItemReq.Price,
	}

	itemEntity, err := s.itemRepository.InsertItem(item)
	if err != nil {
		return nil, err
	}

	return itemEntity.ToItemModel(), nil
}

func (s *adminServiceImpl) EditItem(itemID uint64, updateItemReq *_itemModel.EditItemReq) (*_itemModel.Item, error) {
	updateItemDto := &_itemEntity.UpdateItemDto{
		Name:        updateItemReq.Name,
		Description: updateItemReq.Description,
		Picture:     updateItemReq.Picture,
		Price:       updateItemReq.Price,
	}

	itemEntity, err := s.itemRepository.UpdateItem(itemID, updateItemDto)
	if err != nil {
		return nil, err
	}

	return itemEntity.ToItemModel(), nil
}

func (s *adminServiceImpl) ArchiveItem(itemID uint64) error {
	return s.itemRepository.ArchiveItem(itemID)
}
