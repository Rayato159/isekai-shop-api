package controller

import "github.com/labstack/echo/v4"

type PlayerController interface {
	PlayerProfiling(pctx echo.Context) error
	PlayerProfileEditing(pctx echo.Context) error
	PlayerInventoryListing(pctx echo.Context) error
}
