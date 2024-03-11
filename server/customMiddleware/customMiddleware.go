package customMiddleware

import "github.com/labstack/echo/v4"

type CustomMiddleware interface {
	PlayerAuthorizing(next echo.HandlerFunc) echo.HandlerFunc
	AdminAuthorizing(next echo.HandlerFunc) echo.HandlerFunc
}
