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
