package server

import (
	_oauth2Controller "github.com/Rayato159/isekai-shop-api/modules/oauth2/controller"
	_oauth2Repository "github.com/Rayato159/isekai-shop-api/modules/oauth2/repository"
	_oauth2Service "github.com/Rayato159/isekai-shop-api/modules/oauth2/service"

	"github.com/Rayato159/isekai-shop-api/packages/state"
)

func (s *echoServer) initOauth2Router() {
	router := s.baseRouter.Group("/oauth2/google")

	oauth2Config := s.conf.Oauth2Config
	stateConfig := s.conf.StateConfig

	jwtState := state.NewJwtState(
		stateConfig.Secret,
		stateConfig.ExpiresAt,
		stateConfig.Issuer,
	)

	oauth2Repository := _oauth2Repository.NewGoogleOauth2Repository(s.db)

	oauth2Service := _oauth2Service.NewGoogleOauth2Service(oauth2Repository)

	controller := _oauth2Controller.NewGoogleOauth2Controller(
		oauth2Config,
		s.app.Logger,
		jwtState,
		oauth2Service,
	)

	router.GET("/login", controller.Login)
	router.GET("/login/callback", controller.LoginCallback)
}
