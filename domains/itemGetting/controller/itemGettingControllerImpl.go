package controller

import (
	"net/http"

	_itemGettingException "github.com/Rayato159/isekai-shop-api/domains/itemGetting/exception"
	_itemGettingModel "github.com/Rayato159/isekai-shop-api/domains/itemGetting/model"
	_itemGettingService "github.com/Rayato159/isekai-shop-api/domains/itemGetting/service"
	"github.com/Rayato159/isekai-shop-api/server/writter"
	"github.com/labstack/echo/v4"
)

type itemControllerImpl struct {
	itemService _itemGettingService.ItemService
	logger      echo.Logger
}

func NewItemControllerImpl(itemGettingService _itemGettingService.ItemService, logger echo.Logger) ItemGettingController {
	return &itemControllerImpl{
		itemService: itemGettingService,
		logger:      logger,
	}
}

func (c *itemControllerImpl) ItemListing(pctx echo.Context) error {
	itemFilter := new(_itemGettingModel.ItemFilter)

	if err := pctx.Bind(itemFilter); err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, &_itemGettingException.ItemListingException{})
	}

	itemListingResult, err := c.itemService.ItemListing(itemFilter)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, itemListingResult)
}
