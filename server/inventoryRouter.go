package server

import (
	_inventoryController "github.com/Rayato159/isekai-shop-api/pkg/inventory/controller"
	_inventoryRepository "github.com/Rayato159/isekai-shop-api/pkg/inventory/repository"
	_inventoryService "github.com/Rayato159/isekai-shop-api/pkg/inventory/service"
	_itemShopRepository "github.com/Rayato159/isekai-shop-api/pkg/itemShop/repository"
)

func (s *echoServer) initInventoryRouter(m *customMiddleware) {
	router := s.app.Group("/v1/inventory-searching")

	itemRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)

	inventoryService := _inventoryService.NewInventoryServiceImpl(
		inventoryRepository,
		itemRepository,
	)

	inventoryController := _inventoryController.NewInventoryControllerImpl(inventoryService, s.app.Logger)

	router.GET("", inventoryController.Listing, m.PlayerAuthorizing)
}
