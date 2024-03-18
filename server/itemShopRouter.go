package server

import (
	_inventoryRepository "github.com/Rayato159/isekai-shop-api/domains/inventory/repository"
	_itemShopController "github.com/Rayato159/isekai-shop-api/domains/itemShop/controller"
	_itemShopRepository "github.com/Rayato159/isekai-shop-api/domains/itemShop/repository"
	_itemShopService "github.com/Rayato159/isekai-shop-api/domains/itemShop/service"
	_playerCoinRepository "github.com/Rayato159/isekai-shop-api/domains/playerCoin/repository"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
)

func (s *echoServer) initItemShopRouter(customMiddleware customMiddleware.CustomMiddleware) {
	router := s.baseRouter.Group("/item-shop")

	balancingRepository := _playerCoinRepository.NewPlayerCoinRepositoryImpl(s.db, s.app.Logger)
	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)
	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)

	itemShopService := _itemShopService.NewItemShopServiceImpl(
		itemShopRepository,
		balancingRepository,
		inventoryRepository,
	)

	itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService, s.app.Logger)

	router.GET("", itemShopController.Listing)
	router.POST("/buying", itemShopController.Buying, customMiddleware.PlayerAuthorizing)
	router.POST("/selling", itemShopController.Selling, customMiddleware.PlayerAuthorizing)
}
