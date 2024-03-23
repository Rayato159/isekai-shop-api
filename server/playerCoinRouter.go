package server

import (
	_playerCoinController "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/controller"
	_playerCoinRepository "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/repository"
	_playerCoinService "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/service"
)

func (s *echoServer) initPlayerCoinRouter(m *customMiddleware) {
	router := s.app.Group("/v1/player-coin")

	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepositoryImpl(s.db, s.app.Logger)

	playerCoinService := _playerCoinService.NewPlayerCoinServiceImpl(
		playerCoinRepository,
	)
	playerCoinController := _playerCoinController.NewPlayerCoinControllerImpl(playerCoinService)

	router.POST("", playerCoinController.CoinAdding, m.PlayerAuthorizing)
	router.GET("", playerCoinController.PlayerCoinShowing, m.PlayerAuthorizing)
}
