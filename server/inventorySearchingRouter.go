package server

import (
	_inventoryController "github.com/Rayato159/isekai-shop-api/domains/inventory/controller"
	_inventoryRepository "github.com/Rayato159/isekai-shop-api/domains/inventory/repository"
	_inventoryService "github.com/Rayato159/isekai-shop-api/domains/inventory/service"
	_itemGettingRepository "github.com/Rayato159/isekai-shop-api/domains/itemGetting/repository"
	_playerRepository "github.com/Rayato159/isekai-shop-api/domains/player/repository"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
)

func (s *echoServer) initInventoryRouter(customMiddleware customMiddleware.CustomMiddleware) {
	router := s.baseRouter.Group("/inventory-searching")

	itemRepository := _itemGettingRepository.NewItemGettingRepositoryImpl(s.db, s.app.Logger)
	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)
	playerRepository := _playerRepository.NewPlayerRepositoryImpl(s.db, s.app.Logger)

	inventoryService := _inventoryService.NewInventoryServiceImpl(
		playerRepository,
		inventoryRepository,
		itemRepository,
	)

	inventoryController := _inventoryController.NewInventoryControllerImpl(inventoryService, s.app.Logger)

	router.GET("", inventoryController.Listing, customMiddleware.PlayerAuthorizing)
}
