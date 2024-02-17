package controller

import "github.com/labstack/echo/v4"

type InventoryController interface {
	InventoryListing(pctx echo.Context) error
}
