package controller

import "github.com/labstack/echo/v4"

type Oauth2Controller interface {
	Login(pctx echo.Context) error
	LoginCallback(pctx echo.Context) error
}
