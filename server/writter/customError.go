package writter

import "github.com/labstack/echo/v4"

type (
	CustomErrorResponse interface {
		GetErrorMessage() string
		GetStatusCode() int
	}

	ErrorMessage struct {
		Message string `json:"error"`
	}
)

func Error(c echo.Context, customError CustomErrorResponse) error {
	return c.JSON(
		customError.GetStatusCode(),
		&ErrorMessage{Message: customError.GetErrorMessage()},
	)
}
