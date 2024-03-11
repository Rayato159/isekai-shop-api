package controller

import (
	"net/http"

	_playerModel "github.com/Rayato159/isekai-shop-api/modules/player/model"
	_playerService "github.com/Rayato159/isekai-shop-api/modules/player/service"
	"github.com/Rayato159/isekai-shop-api/modules/utils"
	"github.com/Rayato159/isekai-shop-api/server/writter"
	"github.com/labstack/echo/v4"
)

type playerControllerImpl struct {
	playerService _playerService.PlayerService
	logger        echo.Logger
}

func NewPlayerControllerImpl(
	playerService _playerService.PlayerService,
	logger echo.Logger,
) PlayerController {
	return &playerControllerImpl{
		playerService: playerService,
		logger:        logger,
	}
}

func (c *playerControllerImpl) PlayerProfiling(pctx echo.Context) error {
	playerID, err := utils.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusUnauthorized, err)
	}

	player, err := c.playerService.PlayerProfiling(playerID)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, player)
}

func (c *playerControllerImpl) PlayerProfileEditing(pctx echo.Context) error {
	playerID, err := utils.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	editPlayerReq := new(_playerModel.PlayerProfileEditingReq)
	if err := pctx.Bind(editPlayerReq); err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	player, err := c.playerService.PlayerProfileEditing(playerID, editPlayerReq)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, player)
}
