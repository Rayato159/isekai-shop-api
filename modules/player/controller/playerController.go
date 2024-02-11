package controller

import "github.com/labstack/echo/v4"

type PlayerController interface {
	EditProfile(pctx echo.Context) error
}
