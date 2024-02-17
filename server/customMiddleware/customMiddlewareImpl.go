package customMiddleware

import (
	_oauth2Controller "github.com/Rayato159/isekai-shop-api/modules/oauth2/controller"

	"github.com/Rayato159/isekai-shop-api/config"
	"github.com/labstack/echo/v4"
)

type customMiddlewaresImpl struct {
	oauth2Conf       *config.OAuth2Config
	logger           echo.Logger
	oauth2Controller _oauth2Controller.OAuth2Controller
}

func NewCustomMiddlewaresImpl(
	oauth2Controller _oauth2Controller.OAuth2Controller,
	oauth2Conf *config.OAuth2Config,
	logger echo.Logger,
) CustomMiddleware {
	return &customMiddlewaresImpl{
		oauth2Controller: oauth2Controller,
		oauth2Conf:       oauth2Conf,
		logger:           logger,
	}
}

func (m *customMiddlewaresImpl) PlayerAuthorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(pctx echo.Context) error {
		return m.oauth2Controller.PlayerAuthorize(pctx, next)
	}
}
