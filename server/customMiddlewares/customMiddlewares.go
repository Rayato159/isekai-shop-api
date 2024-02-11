package customMiddlewares

import "github.com/labstack/echo/v4"

type CustomMiddleware interface {
	Authorize(next echo.HandlerFunc) echo.HandlerFunc
}
