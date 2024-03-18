package controller

import (
	"net/http"

	"github.com/Rayato159/isekai-shop-api/domains/common"
	_playerBalancingModel "github.com/Rayato159/isekai-shop-api/domains/playerBalancing/model"
	_playerBalancingService "github.com/Rayato159/isekai-shop-api/domains/playerBalancing/service"
	"github.com/Rayato159/isekai-shop-api/server/writter"
	"github.com/labstack/echo/v4"
)

type playerBalacingControllerImpl struct {
	playerBalacingService _playerBalancingService.PlayerBalancingService
	logger                echo.Logger
}

func NewBalancingControllerImpl(playerBalacingService _playerBalancingService.PlayerBalancingService, logger echo.Logger) BalancingController {
	return &playerBalacingControllerImpl{
		playerBalacingService: playerBalacingService,
		logger:                logger,
	}
}

func (c *playerBalacingControllerImpl) TopUp(pctx echo.Context) error {
	playerID, err := common.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	topUpReq := new(_playerBalancingModel.TopUpReq)

	if err := pctx.Bind(topUpReq); err != nil {
		c.logger.Error("Failed to bind top up request", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}
	topUpReq.PlayerID = playerID

	playerBalacing, err := c.playerBalacingService.TopUp(topUpReq)
	if err != nil {
		c.logger.Error("Failed to top up", err.Error())
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusCreated, playerBalacing)
}

func (c *playerBalacingControllerImpl) PlayerBalanceShowing(pctx echo.Context) error {
	playerID, err := common.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	balance := c.playerBalacingService.PlayerBalanceShowing(playerID)

	return pctx.JSON(http.StatusOK, balance)
}
