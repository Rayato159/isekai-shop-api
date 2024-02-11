package controller

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	_oauth2Exception "github.com/Rayato159/isekai-shop-api/modules/oauth2/exception"
	_oauth2Model "github.com/Rayato159/isekai-shop-api/modules/oauth2/model"
	_oauth2Service "github.com/Rayato159/isekai-shop-api/modules/oauth2/service"
	"github.com/Rayato159/isekai-shop-api/server/writter"

	"github.com/Rayato159/isekai-shop-api/config"
	"github.com/Rayato159/isekai-shop-api/packages/state"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type (
	googleOAuth2Controller struct {
		oauth2Conf        *oauth2.Config
		oauth2UserInfoUrl string
		stateProvider     state.State
		logger            echo.Logger
		oauth2Service     _oauth2Service.OAuth2Service
	}

	oauth2CallbackResponse struct {
		Message string `json:"message"`
	}
)

func NewGoogleOAuth2Controller(
	oauth2Conf *config.OAuth2Config,
	logger echo.Logger,
	stateProvider state.State,
	oauth2Service _oauth2Service.OAuth2Service,
) OAuth2Controller {
	conf := &oauth2.Config{
		ClientID:     oauth2Conf.ClientId,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.RedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint:     google.Endpoint,
	}

	return &googleOAuth2Controller{
		oauth2Conf:        conf,
		oauth2UserInfoUrl: oauth2Conf.UserInfoUrl,
		stateProvider:     stateProvider,
		logger:            logger,
		oauth2Service:     oauth2Service,
	}
}

func (c *googleOAuth2Controller) Login(pctx echo.Context) error {
	state, err := c.stateProvider.GenerateRandomState()
	if err != nil {
		c.logger.Errorf("Error generating state: %s", err.Error())
		return err
	}

	return pctx.Redirect(302, c.oauth2Conf.AuthCodeURL(state))
}

func (c *googleOAuth2Controller) LoginCallback(pctx echo.Context) error {
	ctx := context.Background()

	if err := c.callbackValidate(pctx); err != nil {
		c.logger.Errorf("Error validating callback: %s", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, &_oauth2Exception.Oauth2Exception{})
	}

	token, err := c.oauth2Conf.Exchange(ctx, pctx.QueryParam("code"))
	if err != nil {
		c.logger.Errorf("Error exchanging code for token: %s", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, &_oauth2Exception.Oauth2Exception{})
	}

	client := c.oauth2Conf.Client(ctx, token)

	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Error reading user info: %s", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, &_oauth2Exception.Oauth2Exception{})

	}

	if err := c.oauth2Service.ManageUserAccount(&_oauth2Model.CreateUserInfo{
		Id:      userInfo.Id,
		Email:   userInfo.Email,
		Name:    userInfo.Name,
		Picture: userInfo.Picture,
		UserPassport: &_oauth2Model.UserPassport{
			RefreshToken: token.RefreshToken,
		},
	}); err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, &_oauth2Exception.Oauth2Exception{})
	}

	c.setSameSiteCookie(pctx, "_oauth2_access_token", token.AccessToken)
	c.setSameSiteCookie(pctx, "_oauth2_refresh_token", token.RefreshToken)
	c.setCookie(pctx, "_user_id", userInfo.Id)

	return pctx.JSON(200, &oauth2CallbackResponse{Message: "Login successful"})
}

func (c *googleOAuth2Controller) Logout(pctx echo.Context) error {
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
	resp, err := client.Get(c.oauth2UserInfoUrl)
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

func (c *googleOAuth2Controller) setSameSiteCookie(pctx echo.Context, name, token string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    token,
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
	}

	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) setCookie(pctx echo.Context, name, token string) {
	cookie := &http.Cookie{
		Name:  name,
		Value: token,
		Path:  "/",
	}

	pctx.SetCookie(cookie)
}
