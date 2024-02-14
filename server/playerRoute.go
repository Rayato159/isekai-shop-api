package server

import (
	_playerController "github.com/Rayato159/isekai-shop-api/modules/player/controller"
	_playerRepository "github.com/Rayato159/isekai-shop-api/modules/player/repository"
	_playerService "github.com/Rayato159/isekai-shop-api/modules/player/service"
	"github.com/Rayato159/isekai-shop-api/server/customMiddlewares"
)

func (s *echoServer) initPlayerRouter(middlewares customMiddlewares.CustomMiddleware) {
	router := s.baseRouter.Group("/player")

	playerRepository := _playerRepository.NewPlayerRepositoryImpl(s.db, s.app.Logger)
	playerService := _playerService.NewPlayerServiceImpl(playerRepository, s.app.Logger)
	playerController := _playerController.NewPlayerControllerImpl(playerService, s.app.Logger)

	router.GET("", playerController.GetPlayerProfile, middlewares.Authorize)
	router.PATCH("", playerController.EditPlayerProfile, middlewares.Authorize)
}
