package server

import (
	_oauth2Controller "github.com/Rayato159/isekai-shop-api/modules/oauth2/controller"
	_oauth2Service "github.com/Rayato159/isekai-shop-api/modules/oauth2/service"
	_playerRepository "github.com/Rayato159/isekai-shop-api/modules/player/repository"

	"github.com/Rayato159/isekai-shop-api/packages/state"
)

func (s *echoServer) initOAuth2Router() {
	router := s.baseRouter.Group("/oauth2/google")

	stateConfig := s.conf.StateConfig

	stateProvider := state.NewJwtState(
		[]byte(stateConfig.Secret),
		stateConfig.ExpiresAt,
		stateConfig.Issuer,
	)

	playerRepository := _playerRepository.NewPlayerRepositoryImpl(s.db, s.app.Logger)

	oauth2Service := _oauth2Service.NewGoogleOAuth2Service(
		playerRepository,
		s.app.Logger,
	)

	controller := _oauth2Controller.NewGoogleOAuth2Controller(
		oauth2Service,
		s.conf.OAuth2Config,
		stateProvider,
		s.app.Logger,
	)

	router.GET("/login", controller.Login)
	router.GET("/login/callback", controller.LoginCallback)
	router.DELETE("/logout", controller.Logout)
}
