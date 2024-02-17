package controller

import (
	"net/http"

	_inventoryService "github.com/Rayato159/isekai-shop-api/modules/inventory/service"
	"github.com/Rayato159/isekai-shop-api/modules/utils"
	"github.com/Rayato159/isekai-shop-api/server/writter"
	"github.com/labstack/echo/v4"
)

type inventoryControllerImpl struct {
	inventoryService _inventoryService.InventoryService
	logger           echo.Logger
}

func NewInventoryController(inventoryService _inventoryService.InventoryService, logger echo.Logger) InventoryController {
	return &inventoryControllerImpl{
		inventoryService: inventoryService,
		logger:           logger,
	}
}

func (c *inventoryControllerImpl) InventoryListing(pctx echo.Context) error {
	playerID, err := utils.GetPlayerID(pctx)
	if err != nil {
		c.logger.Error("Failed to get playerID", err.Error())
		return writter.CustomError(pctx, http.StatusUnauthorized, err)
	}

	inventories, err := c.inventoryService.InventoryListing(playerID)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, inventories)
}
