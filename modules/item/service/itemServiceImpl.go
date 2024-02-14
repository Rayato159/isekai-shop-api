package service

import (
	_itemRepository "github.com/Rayato159/isekai-shop-api/modules/item/repository"
	"github.com/labstack/echo/v4"
)

type ItemServiceImpl struct {
	itemRepository _itemRepository.PlayerRepository
	logger         echo.Logger
}

func NewItemServiceImpl(itemRepository _itemRepository.PlayerRepository, logger echo.Logger) ItemService {
	return &ItemServiceImpl{
		itemRepository: itemRepository,
		logger:         logger,
	}
}
