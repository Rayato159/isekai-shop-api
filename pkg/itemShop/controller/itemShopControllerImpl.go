package controller

import (
	"net/http"

	custom "github.com/Rayato159/isekai-shop-api/pkg/custom"
	_itemShopModel "github.com/Rayato159/isekai-shop-api/pkg/itemShop/model"
	_itemShopService "github.com/Rayato159/isekai-shop-api/pkg/itemShop/service"
	"github.com/Rayato159/isekai-shop-api/pkg/validation"
	"github.com/labstack/echo/v4"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopControllerImpl(itemShopService _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {
	itemFilter := new(_itemShopModel.ItemFilter)

	validatingContext := validation.NewCustomEchoRequest(pctx)

	if err := validatingContext.Bind(itemFilter); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	itemListingResult, err := c.itemShopService.Listing(itemFilter)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, itemListingResult)
}

func (c *itemShopControllerImpl) Buying(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	buyingReq := new(_itemShopModel.BuyingReq)

	validatingContext := validation.NewCustomEchoRequest(pctx)

	if err := validatingContext.Bind(buyingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}
	buyingReq.PlayerID = playerID

	result, err := c.itemShopService.Buying(buyingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, result)
}

func (c *itemShopControllerImpl) Selling(pctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	sellingReq := new(_itemShopModel.SellingReq)

	validatingContext := validation.NewCustomEchoRequest(pctx)

	if err := validatingContext.Bind(sellingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}
	sellingReq.PlayerID = playerID

	result, err := c.itemShopService.Selling(sellingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, result)
}
