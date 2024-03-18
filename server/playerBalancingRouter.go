package server

import (
	_playerBalancingController "github.com/Rayato159/isekai-shop-api/domains/playerBalancing/controller"
	_playerBalancingRepository "github.com/Rayato159/isekai-shop-api/domains/playerBalancing/repository"
	_playerBalancingService "github.com/Rayato159/isekai-shop-api/domains/playerBalancing/service"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
)

func (s *echoServer) initPlayerBalancingRouter(customMiddleware customMiddleware.CustomMiddleware) {
	router := s.baseRouter.Group("/player-balancing")

	playerBalancingRepository := _playerBalancingRepository.NewPlayerBalancingRepositoryImpl(s.db, s.app.Logger)

	playerBalancingService := _playerBalancingService.NewPlayerBalancingServiceImpl(
		playerBalancingRepository,
	)
	playerBalancingController := _playerBalancingController.NewBalancingControllerImpl(playerBalancingService, s.app.Logger)

	router.POST("", playerBalancingController.TopUp, customMiddleware.PlayerAuthorizing)
	router.GET("", playerBalancingController.PlayerBalanceShowing, customMiddleware.PlayerAuthorizing)
}
