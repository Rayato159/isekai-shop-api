package controller

import "github.com/labstack/echo/v4"

type OrderController interface {
	PlayerOrderListing(pctx echo.Context) error
}
