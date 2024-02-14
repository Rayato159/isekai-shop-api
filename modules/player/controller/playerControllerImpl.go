package controller

import (
	"net/http"

	_playerModel "github.com/Rayato159/isekai-shop-api/modules/player/model"
	_playerService "github.com/Rayato159/isekai-shop-api/modules/player/service"
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

func (c *playerControllerImpl) GetPlayerProfile(pctx echo.Context) error {
	playerID := pctx.Get("playerID").(string)

	playerProfile, err := c.playerService.GetPlayerProfile(playerID)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, playerProfile)
}

func (c *playerControllerImpl) EditPlayerProfile(pctx echo.Context) error {
	playerID := pctx.Get("playerID").(string)

	updatePlayerReq := new(_playerModel.UpdatePlayerProfile)
	if err := pctx.Bind(updatePlayerReq); err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	playerProfile, err := c.playerService.EditPlayerProfile(playerID, updatePlayerReq)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, playerProfile)
}
