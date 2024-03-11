package controller

import (
	"net/http"

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

func (c *playerControllerImpl) PlayerInventoryListing(pctx echo.Context) error {
	playerID, err := utils.GetPlayerID(pctx)
	if err != nil {
		c.logger.Error("Failed to get playerID", err.Error())
		return writter.CustomError(pctx, http.StatusUnauthorized, err)
	}

	inventories, err := c.playerService.PlayerInventoryListing(playerID)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, inventories)
}
