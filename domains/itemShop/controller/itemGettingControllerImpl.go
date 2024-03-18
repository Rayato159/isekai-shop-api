package controller

import (
	"net/http"

	_itemShopException "github.com/Rayato159/isekai-shop-api/domains/itemShop/exception"
	_itemShopModel "github.com/Rayato159/isekai-shop-api/domains/itemShop/model"
	_itemShopService "github.com/Rayato159/isekai-shop-api/domains/itemShop/service"
	"github.com/Rayato159/isekai-shop-api/server/writter"
	"github.com/labstack/echo/v4"
)

type itemControllerImpl struct {
	itemService _itemShopService.ItemService
	logger      echo.Logger
}

func NewItemControllerImpl(itemShopService _itemShopService.ItemService, logger echo.Logger) ItemShopController {
	return &itemControllerImpl{
		itemService: itemShopService,
		logger:      logger,
	}
}

func (c *itemControllerImpl) Listing(pctx echo.Context) error {
	itemFilter := new(_itemShopModel.ItemFilter)

	if err := pctx.Bind(itemFilter); err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, &_itemShopException.ItemListingException{})
	}

	itemListingResult, err := c.itemService.Listing(itemFilter)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, itemListingResult)
}
