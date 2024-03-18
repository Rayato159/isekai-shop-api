package controller

import (
	"net/http"

	"github.com/Rayato159/isekai-shop-api/domains/common"
	_playerCoinModel "github.com/Rayato159/isekai-shop-api/domains/playerCoin/model"
	_playerCoinService "github.com/Rayato159/isekai-shop-api/domains/playerCoin/service"
	"github.com/Rayato159/isekai-shop-api/server/writter"
	"github.com/labstack/echo/v4"
)

type playerCoinControllerImpl struct {
	playerCoinService _playerCoinService.PlayerCoinService
	logger            echo.Logger
}

func NewPlayerCoinControllerImpl(playerCoinService _playerCoinService.PlayerCoinService, logger echo.Logger) PlayerCoinController {
	return &playerCoinControllerImpl{
		playerCoinService: playerCoinService,
		logger:            logger,
	}
}

func (c *playerCoinControllerImpl) BuyingCoin(pctx echo.Context) error {
	playerID, err := common.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	buyingCoinReq := new(_playerCoinModel.BuyingCoinReq)

	if err := pctx.Bind(buyingCoinReq); err != nil {
		c.logger.Error("Failed to bind top up request", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}
	buyingCoinReq.PlayerID = playerID

	playerCoin, err := c.playerCoinService.BuyingCoin(buyingCoinReq)
	if err != nil {
		c.logger.Error("Failed to top up", err.Error())
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusCreated, playerCoin)
}

func (c *playerCoinControllerImpl) PlayerCoinShowing(pctx echo.Context) error {
	playerID, err := common.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	coin := c.playerCoinService.PlayerCoinShowing(playerID)

	return pctx.JSON(http.StatusOK, coin)
}
