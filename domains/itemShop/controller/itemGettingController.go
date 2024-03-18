package controller

import "github.com/labstack/echo/v4"

type ItemShopController interface {
	Listing(pctx echo.Context) error
}
