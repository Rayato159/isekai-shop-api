package controller

import "github.com/labstack/echo/v4"

type ItemManagingController interface {
	ItemCreating(pctx echo.Context) error
	ItemEditing(pctx echo.Context) error
	ItemArchiving(pctx echo.Context) error
}
