package controller

import "github.com/labstack/echo/v4"

type PurchasingController interface {
	ItemBuying(pctx echo.Context) error
	ItemSelling(pctx echo.Context) error
}
