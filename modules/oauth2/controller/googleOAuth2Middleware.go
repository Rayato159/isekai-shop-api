package controller

import (
	"context"
	"net/http"

	_oauth2Exception "github.com/Rayato159/isekai-shop-api/modules/oauth2/exception"
	"github.com/Rayato159/isekai-shop-api/server/writter"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

func (c *googleOAuth2Controller) Authorize(pctx echo.Context, next echo.HandlerFunc) error {
	ctx := context.Background()

	tokenSource, err := c.getToken(pctx)
	if err != nil {
		c.logger.Errorf("Error reading token: %s", err.Error())
		return writter.CustomError(
			pctx, http.StatusUnauthorized,
			&_oauth2Exception.UnAuthorizeException{},
		)

	}

	// Validate the token
	if !tokenSource.Valid() {
		c.logger.Errorf("Token is not valid")

		// Refresh the token
		newToken, err := c.refreshToken(ctx, tokenSource)
		if err != nil {
			c.logger.Errorf("Error refreshing token: %s", err.Error())
			return writter.CustomError(
				pctx, http.StatusUnauthorized,
				&_oauth2Exception.UnAuthorizeException{},
			)
		}

		// Update the token
		c.setSameSiteCookie(pctx, oauth2AccessTokenCookieName, newToken.AccessToken)
		c.setSameSiteCookie(pctx, oauth2RefreshTokenCookieName, newToken.RefreshToken)

		tokenSource = newToken
	}

	// Get user info
	client := googleOAuth2.Client(ctx, tokenSource)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Error reading user info: %s", err.Error())
		return writter.CustomError(pctx, http.StatusUnauthorized, &_oauth2Exception.UnAuthorizeException{})

	}

	pctx.Set("userId", userInfo.ID)

	return next(pctx)
}

func (c *googleOAuth2Controller) refreshToken(ctx context.Context, token *oauth2.Token) (*oauth2.Token, error) {
	newToken, err := googleOAuth2.TokenSource(ctx, token).Token()
	if err != nil {
		c.logger.Errorf("Error refreshing token: %s", err.Error())
		return nil, err
	}

	return newToken, nil
}

func (c *googleOAuth2Controller) getToken(pctx echo.Context) (*oauth2.Token, error) {
	accessToken, err := pctx.Cookie(oauth2AccessTokenCookieName)
	if err != nil {
		c.logger.Errorf("Error reading access token: %s", err.Error())
		return nil, err
	}

	refreshToken, err := pctx.Cookie(oauth2AccessTokenCookieName)
	if err != nil {
		c.logger.Errorf("Error reading refresh token: %s", err.Error())
		return nil, err
	}

	return &oauth2.Token{
		AccessToken:  accessToken.Value,
		RefreshToken: refreshToken.Value,
	}, nil
}
