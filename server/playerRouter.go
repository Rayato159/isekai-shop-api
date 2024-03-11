package server

import (
	_itemRepository "github.com/Rayato159/isekai-shop-api/modules/item/repository"
	_playerController "github.com/Rayato159/isekai-shop-api/modules/player/controller"
	_playerSource "github.com/Rayato159/isekai-shop-api/modules/player/repository"
	_playerService "github.com/Rayato159/isekai-shop-api/modules/player/service"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
)

func (s *echoServer) initPlayerRouter(customMiddleware customMiddleware.CustomMiddleware) {
	router := s.baseRouter.Group("/player")

	itemRepository := _itemRepository.NewItemRepositoryImpl(s.db, s.app.Logger)
	inventoryRepository := _playerSource.NewInventoryRepositoryImpl(s.db, s.app.Logger)
	playerRepository := _playerSource.NewPlayerRepositoryImpl(s.db, s.app.Logger)
	playerService := _playerService.NewPlayerServiceImpl(playerRepository, inventoryRepository, itemRepository)
	playerController := _playerController.NewPlayerControllerImpl(playerService, s.app.Logger)

	router.GET("", playerController.PlayerProfiling, customMiddleware.PlayerAuthorizing)
	router.PATCH("", playerController.PlayerProfileEditing, customMiddleware.PlayerAuthorizing)
	router.GET("/inventory", playerController.PlayerInventoryListing, customMiddleware.PlayerAuthorizing)
}
