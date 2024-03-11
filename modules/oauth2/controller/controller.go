package controller

import "github.com/labstack/echo/v4"

type OAuth2Controller interface {
	PlayerLogin(pctx echo.Context) error
	AdminLogin(pctx echo.Context) error
	PlayerLoginCallback(pctx echo.Context) error
	AdminLoginCallback(pctx echo.Context) error
	Logout(pctx echo.Context) error

	// Middlewares
	PlayerAuthorizing(pctx echo.Context, next echo.HandlerFunc) error
	AdminAuthorizing(pctx echo.Context, next echo.HandlerFunc) error
}
