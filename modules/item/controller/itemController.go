package controller

import "github.com/labstack/echo/v4"

type ItemController interface {
	ItemListing(pctx echo.Context) error
}
