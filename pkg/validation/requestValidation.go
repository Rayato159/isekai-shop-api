package validation

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type (
	EchoRequest interface {
		Bind(obj any) error
	}

	customEchoRequest struct {
		ctx       echo.Context
		validator *validator.Validate
	}
)

var (
	once              sync.Once
	validatorInstance *validator.Validate
)

func NewCustomEchoRequest(echoRequest echo.Context) EchoRequest {
	once.Do(func() {
		validatorInstance = validator.New()
	})

	return &customEchoRequest{
		ctx:       echoRequest,
		validator: validatorInstance,
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
