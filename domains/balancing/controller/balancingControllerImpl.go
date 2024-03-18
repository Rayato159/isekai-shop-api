package controller

import (
	"net/http"

	_balancingModel "github.com/Rayato159/isekai-shop-api/domains/balancing/model"
	_balancingService "github.com/Rayato159/isekai-shop-api/domains/balancing/service"
	"github.com/Rayato159/isekai-shop-api/domains/utils"
	"github.com/Rayato159/isekai-shop-api/server/writter"
	"github.com/labstack/echo/v4"
)

type balancingControllerImpl struct {
	balancingService _balancingService.BalancingService
	logger           echo.Logger
}

func NewBalancingControllerImpl(balancingService _balancingService.BalancingService, logger echo.Logger) BalancingController {
	return &balancingControllerImpl{
		balancingService: balancingService,
		logger:           logger,
	}
}

func (c *balancingControllerImpl) TopUp(pctx echo.Context) error {
	playerID, err := utils.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	topUpReq := new(_balancingModel.TopUpReq)

	if err := pctx.Bind(topUpReq); err != nil {
		c.logger.Error("Failed to bind top up request", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}
	topUpReq.PlayerID = playerID

	balancing, err := c.balancingService.TopUp(topUpReq)
	if err != nil {
		c.logger.Error("Failed to top up", err.Error())
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusCreated, balancing)
}

func (c *balancingControllerImpl) PlayerBalanceShowing(pctx echo.Context) error {
	playerID, err := utils.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	balance := c.balancingService.PlayerBalanceShowing(playerID)

	return pctx.JSON(http.StatusOK, balance)
}

func (c *balancingControllerImpl) ItemBuying(pctx echo.Context) error {
	playerID, err := utils.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	itemBuyingReq := new(_balancingModel.ItemBuyingReq)

	if err := pctx.Bind(itemBuyingReq); err != nil {
		c.logger.Error("Failed to bind buy item request", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}
	itemBuyingReq.PlayerID = playerID

	balancing, err := c.balancingService.ItemBuying(itemBuyingReq)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, balancing)
}

func (c *balancingControllerImpl) ItemSelling(pctx echo.Context) error {
	playerID, err := utils.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	itemSellingReq := new(_balancingModel.ItemSellingReq)

	if err := pctx.Bind(itemSellingReq); err != nil {
		c.logger.Error("Failed to bind sell item request", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}
	itemSellingReq.PlayerID = playerID

	balancing, err := c.balancingService.ItemSelling(itemSellingReq)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, balancing)
}
