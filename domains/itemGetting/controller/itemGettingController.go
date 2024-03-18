package controller

import "github.com/labstack/echo/v4"

type ItemGettingController interface {
	ItemListing(pctx echo.Context) error
}
