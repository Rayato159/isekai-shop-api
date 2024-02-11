package controller

import "github.com/labstack/echo/v4"

type OAuth2Controller interface {
	Login(pctx echo.Context) error
	LoginCallback(pctx echo.Context) error
	Logout(pctx echo.Context) error
}
