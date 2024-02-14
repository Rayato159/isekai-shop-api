package controller

import "github.com/labstack/echo/v4"

type PlayerController interface {
	GetPlayerProfile(pctx echo.Context) error
	EditPlayerProfile(pctx echo.Context) error
}
