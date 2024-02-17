package controller

import "github.com/labstack/echo/v4"

type PlayerController interface {
	GetPlayer(pctx echo.Context) error
	EditPlayer(pctx echo.Context) error
}
