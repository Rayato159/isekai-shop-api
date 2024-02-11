package server

import (
	_oauth2Controller "github.com/Rayato159/isekai-shop-api/modules/oauth2/controller"
	_oauth2Repository "github.com/Rayato159/isekai-shop-api/modules/oauth2/repository"
	_oauth2Service "github.com/Rayato159/isekai-shop-api/modules/oauth2/service"
	_playerRepository "github.com/Rayato159/isekai-shop-api/modules/player/repository"

	"github.com/Rayato159/isekai-shop-api/packages/state"
)

func (s *echoServer) initOAuth2Router() {
	router := s.baseRouter.Group("/oauth2/google")

	oauth2Config := s.conf.OAuth2Config
	stateConfig := s.conf.StateConfig

	stateProvider := state.NewJwtState(
		stateConfig.Secret,
		stateConfig.ExpiresAt,
		stateConfig.Issuer,
	)

	oauth2Repository := _oauth2Repository.NewGoogleOAuth2Repository(s.db, s.app.Logger)
	playerRepository := _playerRepository.NewPlayerRepositoryImpl(s.db, s.app.Logger)

	oauth2Service := _oauth2Service.NewGoogleOAuth2Service(
		oauth2Repository,
		playerRepository,
	)

	controller := _oauth2Controller.NewGoogleOAuth2Controller(
		oauth2Config,
		s.app.Logger,
		stateProvider,
		oauth2Service,
	)

	router.GET("/login", controller.Login)
	router.GET("/login/callback", controller.LoginCallback)
	router.GET("/logout", controller.Logout)
}
