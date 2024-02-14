package controller

import (
	"net/http"

	_itemException "github.com/Rayato159/isekai-shop-api/modules/item/exception"
	_itemModel "github.com/Rayato159/isekai-shop-api/modules/item/model"
	_itemService "github.com/Rayato159/isekai-shop-api/modules/item/service"
	"github.com/Rayato159/isekai-shop-api/server/writter"
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

func (c *itemControllerImpl) ItemListing(pctx echo.Context) error {
	itemFilter := new(_itemModel.ItemFilter)

	if err := pctx.Bind(itemFilter); err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, &_itemException.ItemListingException{})
	}

	itemListingResult, err := c.itemService.ItemListing(itemFilter)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, itemListingResult)
}
