package controller

import (
	_itemService "github.com/Rayato159/isekai-shop-api/modules/item/service"
	"github.com/labstack/echo/v4"
)

type itemControllerImpl struct {
	itemService _itemService.ItemService
	logger      echo.Logger
}

func NewItemControllerImpl(itemService _itemService.ItemService, logger echo.Logger) ItemController {
	return &itemControllerImpl{
		itemService: itemService,
		logger:      logger,
	}
}
