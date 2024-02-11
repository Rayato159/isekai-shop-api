package customMiddlewares

import (
	"context"
	"net/http"
	"strings"

	_oauth2Exception "github.com/Rayato159/isekai-shop-api/modules/oauth2/exception"
	_oauth2Service "github.com/Rayato159/isekai-shop-api/modules/oauth2/service"
	"github.com/Rayato159/isekai-shop-api/server/writter"

	"github.com/Rayato159/isekai-shop-api/config"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type customMiddlewaresImpl struct {
	oauth2Conf    *oauth2.Config
	logger        echo.Logger
	oauth2Service _oauth2Service.OAuth2Service
}

func NewCustomMiddlewaresImpl(
	oauth2Conf *config.OAuth2Config,
	logger echo.Logger,
	oauth2Service _oauth2Service.OAuth2Service,
) CustomMiddleware {
	conf := &oauth2.Config{
		ClientID:     oauth2Conf.ClientId,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.RedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint:     google.Endpoint,
	}

	return &customMiddlewaresImpl{
		oauth2Conf:    conf,
		logger:        logger,
		oauth2Service: oauth2Service,
	}
}

func (c *customMiddlewaresImpl) Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(pctx echo.Context) error {
		ctx := context.Background()

		accessToken := strings.TrimPrefix(pctx.Request().Header.Get("Authorization"), "Bearer ")

		oauth2Token := &oauth2.Token{
			AccessToken: accessToken,
		}

		// Validate the token
		_, err := c.oauth2Conf.TokenSource(ctx, oauth2Token).Token()

		if err != nil {
			c.logger.Errorf("Error validating token: %s", err.Error())
			return writter.CustomError(
				pctx, http.StatusInternalServerError,
				&_oauth2Exception.UnAuthorizeException{},
			)
		}

		return next(pctx)
	}
}
