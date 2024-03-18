package controller

import (
	"net/http"

	"github.com/Rayato159/isekai-shop-api/domains/controllerUtils"
	"github.com/Rayato159/isekai-shop-api/server/writter"
	"github.com/labstack/echo/v4"

	_purchasingModel "github.com/Rayato159/isekai-shop-api/domains/purchasing/model"
	_purchasingService "github.com/Rayato159/isekai-shop-api/domains/purchasing/service"
)

type purchasingControllerImpl struct {
	purchasingService _purchasingService.PurchasingService
	logger            echo.Logger
}

func NewPurchasingControllerImpl(
	purchasingService _purchasingService.PurchasingService,
	logger echo.Logger,
) PurchasingController {
	return &purchasingControllerImpl{
		purchasingService,
		logger,
	}
}

func (c *purchasingControllerImpl) ItemBuying(pctx echo.Context) error {
	playerID, err := controllerUtils.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	itemBuyingReq := new(_purchasingModel.ItemBuyingReq)

	if err := pctx.Bind(itemBuyingReq); err != nil {
		c.logger.Error("Failed to bind buy item request", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}
	itemBuyingReq.PlayerID = playerID

	purchasing, err := c.purchasingService.ItemBuying(itemBuyingReq)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, purchasing)
}

func (c *purchasingControllerImpl) ItemSelling(pctx echo.Context) error {
	playerID, err := controllerUtils.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	itemSellingReq := new(_purchasingModel.ItemSellingReq)

	if err := pctx.Bind(itemSellingReq); err != nil {
		c.logger.Error("Failed to bind sell item request", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}
	itemSellingReq.PlayerID = playerID

	purchasing, err := c.purchasingService.ItemSelling(itemSellingReq)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, purchasing)
}
