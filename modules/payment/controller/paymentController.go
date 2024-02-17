package controller

import "github.com/labstack/echo/v4"

type PaymentController interface {
	TopUp(pctx echo.Context) error
	CalculatePlayerBalance(pctx echo.Context) error
}
