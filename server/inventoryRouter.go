package server

import (
	_inventoryController "github.com/Rayato159/isekai-shop-api/modules/inventory/controller"
	_inventoryRepository "github.com/Rayato159/isekai-shop-api/modules/inventory/repository"
	_inventoryService "github.com/Rayato159/isekai-shop-api/modules/inventory/service"
	_itemRepository "github.com/Rayato159/isekai-shop-api/modules/item/repository"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
)

func (s *echoServer) initInventoryRouter(customMiddleware customMiddleware.CustomMiddleware) {
	router := s.baseRouter.Group("/inventory")

	itemRepository := _itemRepository.NewItemRepositoryImpl(s.db, s.app.Logger)
	inventoryRepository := _inventoryRepository.NewInventoryRepository(s.db, s.app.Logger)
	inventoryService := _inventoryService.NewInventoryService(inventoryRepository, itemRepository, s.app.Logger)
	itemController := _inventoryController.NewInventoryController(inventoryService, s.app.Logger)

	router.GET("", itemController.InventoryListing, customMiddleware.PlayerAuthorize)
}
