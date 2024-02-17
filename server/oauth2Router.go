package server

import (
	_adminRepository "github.com/Rayato159/isekai-shop-api/modules/admin/repository"
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
	adminRepository := _adminRepository.NewAdminRepositoryImpl(s.db, s.app.Logger)

	oauth2Service := _oauth2Service.NewGoogleOAuth2Service(
		playerRepository,
		adminRepository,
	)

	oauth2Controller := _oauth2Controller.NewGoogleOAuth2Controller(
		oauth2Service,
		s.conf.OAuth2Config,
		stateProvider,
		s.app.Logger,
	)

	router.GET("/player/login", oauth2Controller.PlayerLogin)
	router.GET("/admin/login", oauth2Controller.AdminLogin)
	router.GET("/player/login/callback", oauth2Controller.PlayerLoginCallback)
	router.GET("/admin/login/callback", oauth2Controller.AdminLoginCallback)
	router.DELETE("/logout", oauth2Controller.Logout)
}
