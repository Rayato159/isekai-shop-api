package server

import (
	_inventoryRepository "github.com/Rayato159/isekai-shop-api/pkg/inventory/repository"
	_itemShopController "github.com/Rayato159/isekai-shop-api/pkg/itemShop/controller"
	_itemShopRepository "github.com/Rayato159/isekai-shop-api/pkg/itemShop/repository"
	_itemShopService "github.com/Rayato159/isekai-shop-api/pkg/itemShop/service"
	_playerCoinRepository "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/repository"
)

func (s *echoServer) initItemShopRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/item-shop")

	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepositoryImpl(s.db, s.app.Logger)
	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)
	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)

	itemShopService := _itemShopService.NewItemShopServiceImpl(
		itemShopRepository,
		playerCoinRepository,
		inventoryRepository,
		s.app.Logger,
	)

	itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router.GET("", itemShopController.Listing)
	router.POST("/buying", itemShopController.Buying, m.PlayerAuthorizing)
	router.POST("/selling", itemShopController.Selling, m.PlayerAuthorizing)
}
