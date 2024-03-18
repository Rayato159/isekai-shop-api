package server

import (
	_playerCoinController "github.com/Rayato159/isekai-shop-api/domains/playerCoin/controller"
	_playerCoinRepository "github.com/Rayato159/isekai-shop-api/domains/playerCoin/repository"
	_playerCoinService "github.com/Rayato159/isekai-shop-api/domains/playerCoin/service"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
)

func (s *echoServer) initPlayerCoinRouter(customMiddleware customMiddleware.CustomMiddleware) {
	router := s.baseRouter.Group("/player-coin")

	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepositoryImpl(s.db, s.app.Logger)

	playerCoinService := _playerCoinService.NewPlayerCoinServiceImpl(
		playerCoinRepository,
	)
	playerCoinController := _playerCoinController.NewPlayerCoinControllerImpl(playerCoinService, s.app.Logger)

	router.POST("", playerCoinController.BuyingCoin, customMiddleware.PlayerAuthorizing)
	router.GET("", playerCoinController.PlayerCoinShowing, customMiddleware.PlayerAuthorizing)
}
