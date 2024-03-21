package controller

import (
	"context"
	"net/http"

	"github.com/Rayato159/isekai-shop-api/pkg/custom"
	_oauth2 "github.com/Rayato159/isekai-shop-api/pkg/oauth2/exception"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

func (c *googleOAuth2Controller) PlayerAuthorizing(pctx echo.Context, next echo.HandlerFunc) error {
	ctx := context.Background()

	tokenSource, err := c.getToken(pctx)
	if err != nil {
		c.logger.Errorf("Error reading token: %s", err.Error())
		return custom.CustomError(
			pctx, http.StatusUnauthorized,
			&_oauth2.UnAuthorize{},
		)

	}

	// Validate the token
	if !tokenSource.Valid() {
		c.logger.Errorf("Token is not valid")

		// Refresh the token
		tokenSource, err = c.playerRefreshToken(pctx, tokenSource)
		if err != nil {
			c.logger.Errorf("Error refreshing token: %s", err.Error())
			return custom.CustomError(
				pctx, http.StatusUnauthorized,
				&_oauth2.UnAuthorize{},
			)
		}
	}

	// Get user info
	client := playerGoogleOAuth2.Client(ctx, tokenSource)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Error reading user info: %s", err.Error())
		return custom.CustomError(pctx, http.StatusUnauthorized, &_oauth2.UnAuthorize{})

	}

	if !c.oauth2Service.IsThisGuyIsReallyPlayer(userInfo.ID) {
		return custom.CustomError(pctx, http.StatusUnauthorized, &_oauth2.NoPermission{})
	}

	pctx.Set("playerID", userInfo.ID)

	return next(pctx)
}

func (c *googleOAuth2Controller) AdminAuthorizing(pctx echo.Context, next echo.HandlerFunc) error {
	ctx := context.Background()

	tokenSource, err := c.getToken(pctx)
	if err != nil {
		c.logger.Errorf("Error reading token: %s", err.Error())
		return custom.CustomError(
			pctx, http.StatusUnauthorized,
			&_oauth2.UnAuthorize{},
		)

	}

	// Validate the token
	if !tokenSource.Valid() {
		c.logger.Errorf("Token is not valid")

		// Refresh the token
		tokenSource, err = c.adminRefreshToken(pctx, tokenSource)
		if err != nil {
			c.logger.Errorf("Error refreshing token: %s", err.Error())
			return custom.CustomError(
				pctx, http.StatusUnauthorized,
				&_oauth2.UnAuthorize{},
			)
		}
	}

	// Get user info
	client := adminGoogleOAuth2.Client(ctx, tokenSource)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Error reading user info: %s", err.Error())
		return custom.CustomError(pctx, http.StatusUnauthorized, &_oauth2.UnAuthorize{})

	}

	if !c.oauth2Service.IsThisGuyIsReallyAdmin(userInfo.ID) {
		return custom.CustomError(pctx, http.StatusUnauthorized, &_oauth2.NoPermission{})
	}

	pctx.Set("adminID", userInfo.ID)

	return next(pctx)
}

func (c *googleOAuth2Controller) playerRefreshToken(pctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	ctx := pctx.Request().Context()

	updatedToken, err := playerGoogleOAuth2.TokenSource(ctx, token).Token()
	if err != nil {
		c.logger.Errorf("Error refreshing token: %s", err.Error())
		return nil, err
	}

	// Update cookies
	c.setSameSiteCookie(pctx, oauth2AccessTokenCookieName, updatedToken.AccessToken)
	c.setSameSiteCookie(pctx, oauth2RefreshTokenCookieName, updatedToken.RefreshToken)

	return updatedToken, nil
}

func (c *googleOAuth2Controller) adminRefreshToken(pctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	ctx := pctx.Request().Context()

	updatedToken, err := adminGoogleOAuth2.TokenSource(ctx, token).Token()
	if err != nil {
		c.logger.Errorf("Error refreshing token: %s", err.Error())
		return nil, err
	}

	// Update cookies
	c.setSameSiteCookie(pctx, oauth2AccessTokenCookieName, updatedToken.AccessToken)
	c.setSameSiteCookie(pctx, oauth2RefreshTokenCookieName, updatedToken.RefreshToken)

	return updatedToken, nil
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
