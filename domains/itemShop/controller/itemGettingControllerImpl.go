package controller

import (
	"net/http"

	"github.com/Rayato159/isekai-shop-api/domains/common"
	_itemShopException "github.com/Rayato159/isekai-shop-api/domains/itemShop/exception"
	_itemShopModel "github.com/Rayato159/isekai-shop-api/domains/itemShop/model"
	_itemShopService "github.com/Rayato159/isekai-shop-api/domains/itemShop/service"
	"github.com/Rayato159/isekai-shop-api/server/writter"
	"github.com/labstack/echo/v4"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
	logger          echo.Logger
}

func NewItemShopControllerImpl(itemShopService _itemShopService.ItemShopService, logger echo.Logger) ItemShopController {
	return &itemShopControllerImpl{
		itemShopService,
		logger,
	}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {
	itemFilter := new(_itemShopModel.ItemFilter)

	if err := pctx.Bind(itemFilter); err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, &_itemShopException.ItemListingException{})
	}

	itemListingResult, err := c.itemShopService.Listing(itemFilter)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, itemListingResult)
}

func (c *itemShopControllerImpl) Buying(pctx echo.Context) error {
	playerID, err := common.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	buyingReq := new(_itemShopModel.BuyingReq)

	if err := pctx.Bind(buyingReq); err != nil {
		c.logger.Error("Failed to bind buy item request", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}
	buyingReq.PlayerID = playerID

	result, err := c.itemShopService.Buying(buyingReq)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, result)
}

func (c *itemShopControllerImpl) Selling(pctx echo.Context) error {
	playerID, err := common.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	sellingReq := new(_itemShopModel.SellingReq)

	if err := pctx.Bind(sellingReq); err != nil {
		c.logger.Error("Failed to bind sell item request", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}
	sellingReq.PlayerID = playerID

	result, err := c.itemShopService.Selling(sellingReq)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, result)
}
