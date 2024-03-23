package custom

import "github.com/labstack/echo/v4"

type ErrorMessage struct {
	Message string `json:"error"`
}

func Error(c echo.Context, statusCode int, err error) error {
	return c.JSON(
		statusCode,
		&ErrorMessage{Message: err.Error()},
	)
}
