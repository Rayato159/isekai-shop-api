package controller

import "github.com/labstack/echo/v4"

type BalancingController interface {
	TopUp(pctx echo.Context) error
	PlayerBalanceShowing(pctx echo.Context) error
	ItemBuying(pctx echo.Context) error
	ItemSelling(pctx echo.Context) error
}
