package controller

import "github.com/labstack/echo/v4"

type PlayerCoinController interface {
	BuyingCoin(pctx echo.Context) error
	PlayerBalanceShowing(pctx echo.Context) error
}
