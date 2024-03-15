package controller

import "github.com/labstack/echo/v4"

type PlayerController interface {
	PlayerInventoryListing(pctx echo.Context) error
}
