package controller

import "github.com/labstack/echo/v4"

type PlayerCoinController interface {
	BuyingCoin(pctx echo.Context) error
	PlayerCoinShowing(pctx echo.Context) error
}
