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
		return custom.Error(
			pctx, http.StatusUnauthorized,
			&_oauth2.Unauthorized{},
		)

	}

	// Validate the token
	if !tokenSource.Valid() {
		c.logger.Errorf("Token is not valid")

		// Refresh the token
		tokenSource, err = c.playerTokenRefreshing(pctx, tokenSource)
		if err != nil {
			c.logger.Errorf("Error refreshing token: %s", err.Error())
			return custom.Error(pctx, http.StatusUnauthorized, err)
		}
	}

	// Get user info
	client := playerGoogleOAuth2.Client(ctx, tokenSource)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Error reading user info: %s", err.Error())
		return custom.Error(pctx, http.StatusUnauthorized, err)

	}

	if !c.oauth2Service.IsThisGuyIsReallyPlayer(userInfo.ID) {
		return custom.Error(pctx, http.StatusUnauthorized, &_oauth2.NoPermission{})
	}

	pctx.Set("playerID", userInfo.ID)

	return next(pctx)
}

func (c *googleOAuth2Controller) AdminAuthorizing(pctx echo.Context, next echo.HandlerFunc) error {
	ctx := context.Background()

	tokenSource, err := c.getToken(pctx)
	if err != nil {
		c.logger.Errorf("Error reading token: %s", err.Error())
		return custom.Error(pctx, http.StatusUnauthorized, err)

	}

	// Validate the token
	if !tokenSource.Valid() {
		c.logger.Errorf("Token is not valid")

		// Refresh the token
		tokenSource, err = c.adminTokenRefreshing(pctx, tokenSource)
		if err != nil {
			c.logger.Errorf("Error refreshing token: %s", err.Error())
			return custom.Error(pctx, http.StatusUnauthorized, err)
		}
	}

	// Get user info
	client := adminGoogleOAuth2.Client(ctx, tokenSource)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Error reading user info: %s", err.Error())
		return custom.Error(pctx, http.StatusUnauthorized, &_oauth2.Unauthorized{})

	}

	if !c.oauth2Service.IsThisGuyIsReallyAdmin(userInfo.ID) {
		return custom.Error(pctx, http.StatusUnauthorized, &_oauth2.NoPermission{})
	}

	pctx.Set("adminID", userInfo.ID)

	return next(pctx)
}

func (c *googleOAuth2Controller) playerTokenRefreshing(pctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	ctx := context.Background()

	updatedToken, err := playerGoogleOAuth2.TokenSource(ctx, token).Token()
	if err != nil {
		c.logger.Errorf("Error refreshing token: %s", err.Error())
		return nil, &_oauth2.Unauthorized{}
	}

	// Update cookies
	c.setSameSiteCookie(pctx, accessTokenCookieName, updatedToken.AccessToken)
	c.setSameSiteCookie(pctx, refreshTokenCookieName, updatedToken.RefreshToken)

	return updatedToken, nil
}

func (c *googleOAuth2Controller) adminTokenRefreshing(pctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	ctx := context.Background()

	updatedToken, err := adminGoogleOAuth2.TokenSource(ctx, token).Token()
	if err != nil {
		c.logger.Errorf("Error refreshing token: %s", err.Error())
		return nil, &_oauth2.Unauthorized{}
	}

	// Update cookies
	c.setSameSiteCookie(pctx, accessTokenCookieName, updatedToken.AccessToken)
	c.setSameSiteCookie(pctx, refreshTokenCookieName, updatedToken.RefreshToken)

	return updatedToken, nil
}

func (c *googleOAuth2Controller) getToken(pctx echo.Context) (*oauth2.Token, error) {
	accessToken, err := pctx.Cookie(accessTokenCookieName)
	if err != nil {
		c.logger.Errorf("Error reading access token: %s", err.Error())
		return nil, &_oauth2.Unauthorized{}
	}

	refreshToken, err := pctx.Cookie(accessTokenCookieName)
	if err != nil {
		c.logger.Errorf("Error reading refresh token: %s", err.Error())
		return nil, &_oauth2.Unauthorized{}
	}

	return &oauth2.Token{
		AccessToken:  accessToken.Value,
		RefreshToken: refreshToken.Value,
	}, nil
}
