package controller

import "github.com/labstack/echo/v4"

type AdminController interface {
	CreateItem(pctx echo.Context) error
	EditItem(pctx echo.Context) error
	ArchiveItem(pctx echo.Context) error
}
