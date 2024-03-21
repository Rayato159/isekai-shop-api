package controller

import "github.com/labstack/echo/v4"

type ItemManagingController interface {
	Creating(pctx echo.Context) error
	Editing(pctx echo.Context) error
	Archiving(pctx echo.Context) error
}
