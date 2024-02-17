package customMiddleware

import "github.com/labstack/echo/v4"

type CustomMiddleware interface {
	PlayerAuthorize(next echo.HandlerFunc) echo.HandlerFunc
	AdminAuthorize(next echo.HandlerFunc) echo.HandlerFunc
}
