package server

import (
	_inventoryRepository "github.com/Rayato159/isekai-shop-api/domains/inventory/repository"
	_itemShopRepository "github.com/Rayato159/isekai-shop-api/domains/itemShop/repository"
	_playerBalancingRepository "github.com/Rayato159/isekai-shop-api/domains/playerBalancing/repository"
	_purchasingController "github.com/Rayato159/isekai-shop-api/domains/purchasing/controller"
	_purchasingRepository "github.com/Rayato159/isekai-shop-api/domains/purchasing/repository"
	_purchasingService "github.com/Rayato159/isekai-shop-api/domains/purchasing/service"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
)

func (s *echoServer) initPurchasingRouter(customMiddleware customMiddleware.CustomMiddleware) {
	router := s.baseRouter.Group("/purchasing")

	balancingRepository := _playerBalancingRepository.NewPlayerBalancingRepositoryImpl(s.db, s.app.Logger)
	itemRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	purchasingRepository := _purchasingRepository.NewPurchasingRepositoryImpl(s.db, s.app.Logger)
	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)

	purchasingService := _purchasingService.NewPurchasingServiceImpl(
		balancingRepository,
		itemRepository,
		purchasingRepository,
		inventoryRepository,
	)
	purchasingController := _purchasingController.NewPurchasingControllerImpl(purchasingService, s.app.Logger)

	router.POST("/buy", purchasingController.ItemBuying, customMiddleware.PlayerAuthorizing)
	router.POST("/sell", purchasingController.ItemSelling, customMiddleware.PlayerAuthorizing)
}
