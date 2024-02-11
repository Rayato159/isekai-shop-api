package writter

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type (
	Request interface {
		Bind(obj any) error
	}

	customEchoRequest struct {
		ctx       echo.Context
		validator *validator.Validate
	}
)

func NewCustomEchoRequest(echoRequest echo.Context) Request {
	return &customEchoRequest{
		ctx:       echoRequest,
		validator: validator.New(),
	}
}

func (r *customEchoRequest) Bind(data any) error {
	if err := r.ctx.Bind(data); err != nil {
		log.Errorf("Error binding request: %s", err.Error())
		return err
	}

	if err := r.validator.Struct(data); err != nil {
		log.Errorf("Error validating request: %s", err.Error())
		return err
	}

	return nil
}
