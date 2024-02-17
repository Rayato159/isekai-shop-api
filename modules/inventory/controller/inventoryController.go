package controller

import "github.com/labstack/echo/v4"

type InventoryController interface {
	PlayerInventoryListing(pctx echo.Context) error
}
