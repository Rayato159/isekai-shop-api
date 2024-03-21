package controller

import (
	"net/http"

	"github.com/Rayato159/isekai-shop-api/pkg/common"
	_playerCoinModel "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/model"
	_playerCoinService "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/service"
	"github.com/Rayato159/isekai-shop-api/server/validation"
	"github.com/Rayato159/isekai-shop-api/server/writter"
	"github.com/labstack/echo/v4"
)

type playerCoinControllerImpl struct {
	playerCoinService _playerCoinService.PlayerCoinService
}

func NewPlayerCoinControllerImpl(playerCoinService _playerCoinService.PlayerCoinService) PlayerCoinController {
	return &playerCoinControllerImpl{
		playerCoinService: playerCoinService,
	}
}

func (c *playerCoinControllerImpl) CoinAdding(pctx echo.Context) error {
	playerID, err := common.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	coinAddingReq := new(_playerCoinModel.CoinAddingReq)

	validatingContext := validation.NewCustomEchoRequest(pctx)

	if err := validatingContext.Bind(coinAddingReq); err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}
	coinAddingReq.PlayerID = playerID

	playerCoin, err := c.playerCoinService.CoinAdding(coinAddingReq)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusCreated, playerCoin)
}

func (c *playerCoinControllerImpl) PlayerCoinShowing(pctx echo.Context) error {
	playerID, err := common.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	playerCoin := c.playerCoinService.PlayerCoinShowing(playerID)

	return pctx.JSON(http.StatusOK, playerCoin)
}
