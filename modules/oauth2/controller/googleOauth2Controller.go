package controller

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	_oauth2Service "github.com/Rayato159/isekai-shop-api/modules/oauth2/service"

	"github.com/Rayato159/isekai-shop-api/config"
	"github.com/Rayato159/isekai-shop-api/packages/state"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type googleOauth2Controller struct {
	oauth2Conf      *oauth2.Config
	oauth2LocalConf *config.Oauth2Config
	stateProvider   state.State
	logger          echo.Logger
	oauth2Service   _oauth2Service.Oauth2Service
}

func NewGoogleOauth2Controller(
	oauth2Conf *config.Oauth2Config,
	logger echo.Logger,
	stateProvider state.State,
	oauth2Service _oauth2Service.Oauth2Service,
) Oauth2Controller {
	conf := &oauth2.Config{
		ClientID:     oauth2Conf.ClientId,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.RedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint:     google.Endpoint,
	}

	return &googleOauth2Controller{
		oauth2Conf:      conf,
		oauth2LocalConf: oauth2Conf,
		stateProvider:   stateProvider,
		logger:          logger,
		oauth2Service:   oauth2Service,
	}
}

func (c *googleOauth2Controller) Login(pctx echo.Context) error {
	state, err := c.stateProvider.GenerateRandomState()
	if err != nil {
		c.logger.Errorf("Error generating state: %s", err.Error())
		return err
	}

	return pctx.Redirect(302, c.oauth2Conf.AuthCodeURL(state))
}

func (c *googleOauth2Controller) LoginCallback(pctx echo.Context) error {
	ctx := context.Background()

	if err := c.callbackValidate(pctx); err != nil {
		return err
	}

	token, err := c.oauth2Conf.Exchange(ctx, pctx.QueryParam("code"))
	if err != nil {
		c.logger.Errorf("Error exchanging code for token: %s", err.Error())
		return err
	}

	client := c.oauth2Conf.Client(ctx, token)
	userInfo, err := c.getUserInfo(client)
	if err != nil {
		c.logger.Errorf("Error reading user info: %s", err.Error())
		return err

	}
	_ = userInfo

	if err := c.oauth2Service.UpdateCredential(); err != nil {
		return err
	}

	return nil
}

func (c *googleOauth2Controller) callbackValidate(pctx echo.Context) error {
	state := pctx.QueryParam("state")

	if err := c.stateProvider.ParseState(state); err != nil {
		c.logger.Errorf("Error parsing state: %s", err.Error())
		return err
	}

	return nil
}

func (c *googleOauth2Controller) getUserInfo(client *http.Client) (any, error) {
	resp, err := client.Get(c.oauth2LocalConf.UserInfoUrl)
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

	var userInfo any
	if err := json.Unmarshal(userInfoInBytes, &userInfo); err != nil {
		c.logger.Errorf("Error unmarshalling user info: %s", err.Error())
		return nil, err
	}

	return userInfo, nil
}
