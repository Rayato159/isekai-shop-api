package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	_adminModel "github.com/Rayato159/isekai-shop-api/pkg/admin/model"
	_oauth2 "github.com/Rayato159/isekai-shop-api/pkg/oauth2/exception"
	_oauth2Model "github.com/Rayato159/isekai-shop-api/pkg/oauth2/model"
	_oauth2Service "github.com/Rayato159/isekai-shop-api/pkg/oauth2/service"
	_playerModel "github.com/Rayato159/isekai-shop-api/pkg/player/model"
	"github.com/Rayato159/isekai-shop-api/server/writter"

	"github.com/Rayato159/isekai-shop-api/config"
	"github.com/Rayato159/isekai-shop-api/packages/state"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type (
	googleOAuth2Controller struct {
		oauth2Service _oauth2Service.OAuth2Service
		oauth2Conf    *config.OAuth2Config
		stateProvider state.State
		logger        echo.Logger
	}
)

var (
	playerGoogleOAuth2 *oauth2.Config
	adminGoogleOAuth2  *oauth2.Config
	once               sync.Once

	oauth2AccessTokenCookieName  = "act"
	oauth2RefreshTokenCookieName = "rft"
)

func NewGoogleOAuth2Controller(
	oauth2Service _oauth2Service.OAuth2Service,
	oauth2Conf *config.OAuth2Config,
	stateProvider state.State,
	logger echo.Logger,
) OAuth2Controller {
	once.Do(func() {
		setGooleOAuth2Config(oauth2Conf)
	})

	return &googleOAuth2Controller{
		oauth2Service: oauth2Service,
		oauth2Conf:    oauth2Conf,
		stateProvider: stateProvider,
		logger:        logger,
	}
}

func setGooleOAuth2Config(oauth2Conf *config.OAuth2Config) {
	playerGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Conf.ClientId,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.PlayerRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.Endpoints.AuthUrl,
			TokenURL:      oauth2Conf.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}

	adminGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Conf.ClientId,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.AdminRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.Endpoints.AuthUrl,
			TokenURL:      oauth2Conf.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}
}

func (c *googleOAuth2Controller) PlayerLogin(pctx echo.Context) error {
	state, err := c.stateProvider.GenerateRandomState()

	if err != nil {
		c.logger.Errorf("Error generating state: %s", err.Error())
		return err
	}

	return pctx.Redirect(http.StatusFound, playerGoogleOAuth2.AuthCodeURL(state))
}

func (c *googleOAuth2Controller) AdminLogin(pctx echo.Context) error {
	state, err := c.stateProvider.GenerateRandomState()

	if err != nil {
		c.logger.Errorf("Error generating state: %s", err.Error())
		return err
	}

	return pctx.Redirect(http.StatusFound, adminGoogleOAuth2.AuthCodeURL(state))
}

func (c *googleOAuth2Controller) PlayerLoginCallback(pctx echo.Context) error {
	ctx := context.Background()

	if err := c.callbackValidate(pctx); err != nil {
		c.logger.Errorf("Error validating callback: %s", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, &_oauth2.OAuth2Processing{})
	}

	token, err := playerGoogleOAuth2.Exchange(ctx, pctx.QueryParam("code"))
	if err != nil {
		c.logger.Errorf("Error exchanging code for token: %s", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, &_oauth2.OAuth2Processing{})
	}

	client := playerGoogleOAuth2.Client(ctx, token)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Error reading user info: %s", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, &_oauth2.OAuth2Processing{})

	}

	playerCreatingReq := &_playerModel.CreatePlayerReq{
		ID:     userInfo.ID,
		Email:  userInfo.Email,
		Name:   userInfo.Name,
		Avatar: userInfo.Picture,
	}

	if err := c.oauth2Service.PlayerAccountCreating(playerCreatingReq); err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, &_oauth2.OAuth2Processing{})
	}

	c.setSameSiteCookie(pctx, oauth2AccessTokenCookieName, token.AccessToken)
	c.setSameSiteCookie(pctx, oauth2RefreshTokenCookieName, token.RefreshToken)

	return pctx.JSON(http.StatusOK, &_oauth2Model.LoginResponse{Message: "Login successful"})
}

func (c *googleOAuth2Controller) AdminLoginCallback(pctx echo.Context) error {
	ctx := context.Background()

	if err := c.callbackValidate(pctx); err != nil {
		c.logger.Errorf("Error validating callback: %s", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, &_oauth2.OAuth2Processing{})
	}

	token, err := adminGoogleOAuth2.Exchange(ctx, pctx.QueryParam("code"))
	if err != nil {
		c.logger.Errorf("Error exchanging code for token: %s", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, &_oauth2.OAuth2Processing{})
	}

	client := adminGoogleOAuth2.Client(ctx, token)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Error reading user info: %s", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, &_oauth2.OAuth2Processing{})

	}

	createAdminReq := &_adminModel.CreateAdminReq{
		ID:     userInfo.ID,
		Email:  userInfo.Email,
		Name:   userInfo.Name,
		Avatar: userInfo.Picture,
	}

	if err := c.oauth2Service.AdminAccountCreating(createAdminReq); err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, &_oauth2.OAuth2Processing{})
	}

	c.setSameSiteCookie(pctx, oauth2AccessTokenCookieName, token.AccessToken)
	c.setSameSiteCookie(pctx, oauth2RefreshTokenCookieName, token.RefreshToken)

	return pctx.JSON(http.StatusOK, &_oauth2Model.LoginResponse{Message: "Login successful"})
}

func (c *googleOAuth2Controller) Logout(pctx echo.Context) error {
	accessToken, err := pctx.Cookie(oauth2AccessTokenCookieName)
	if err != nil {
		c.logger.Errorf("Error reading access token: %s", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, &_oauth2.Logout{})
	}

	if err := c.revokeToken(accessToken.Value); err != nil {
		c.logger.Errorf("Error revoking token: %s", err.Error())
		return writter.CustomError(pctx, http.StatusInternalServerError, &_oauth2.Logout{})
	}

	c.removeSameSiteCookie(pctx, oauth2AccessTokenCookieName)
	c.removeSameSiteCookie(pctx, oauth2RefreshTokenCookieName)

	return pctx.JSON(http.StatusOK, &_oauth2Model.LogoutResponse{Message: "Logout successful"})
}

func (c *googleOAuth2Controller) revokeToken(accessToken string) error {
	revokeURL := fmt.Sprintf("%s?token=%s", c.oauth2Conf.RevokeUrl, accessToken)

	resp, err := http.Post(revokeURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		fmt.Println("Error revoking token:", err)
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *googleOAuth2Controller) callbackValidate(pctx echo.Context) error {
	state := pctx.QueryParam("state")

	if err := c.stateProvider.ParseState(state); err != nil {
		c.logger.Errorf("Error parsing state: %s", err.Error())
		return err
	}

	return nil
}

func (c *googleOAuth2Controller) getUserInfo(client *http.Client) (*_oauth2Model.UserInfo, error) {
	resp, err := client.Get(c.oauth2Conf.UserInfoUrl)
	if err != nil {
		c.logger.Errorf("Error getting user info: %s", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	userInfoInBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Errorf("Error reading user info: %s", err.Error())
		return nil, err
	}

	userInfo := new(_oauth2Model.UserInfo)
	if err := json.Unmarshal(userInfoInBytes, &userInfo); err != nil {
		c.logger.Errorf("Error unmarshalling user info: %s", err.Error())
		return nil, err
	}

	return userInfo, nil
}

func (c *googleOAuth2Controller) setSameSiteCookie(pctx echo.Context, name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
	}

	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) removeSameSiteCookie(pctx echo.Context, name string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
	}

	pctx.SetCookie(cookie)
}
