package controller

import (
	"net/http"

	"github.com/Rayato159/isekai-shop-api/domains/common"
	_inventoryService "github.com/Rayato159/isekai-shop-api/domains/inventory/service"
	"github.com/Rayato159/isekai-shop-api/server/writter"
	"github.com/labstack/echo/v4"
)

type inventoryControllerImpl struct {
	inventoryService _inventoryService.InventoryService
	logger           echo.Logger
}

func NewInventoryControllerImpl(
	inventoryService _inventoryService.InventoryService,
	logger echo.Logger,
) InventoryController {
	return &inventoryControllerImpl{
		inventoryService: inventoryService,
		logger:           logger,
	}
}

func (c *inventoryControllerImpl) Listing(pctx echo.Context) error {
	playerID, err := common.GetPlayerID(pctx)
	if err != nil {
		c.logger.Error("Failed to get playerID", err.Error())
		return writter.CustomError(pctx, http.StatusUnauthorized, err)
	}

	inventoryListing, err := c.inventoryService.Listing(playerID)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, inventoryListing)
}
