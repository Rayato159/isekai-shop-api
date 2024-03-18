package server

import (
	_balancingRepository "github.com/Rayato159/isekai-shop-api/domains/balancing/repository"
	_itemGettingRepository "github.com/Rayato159/isekai-shop-api/domains/itemGetting/repository"
	_playerSouce "github.com/Rayato159/isekai-shop-api/domains/player/repository"
	_purchasingController "github.com/Rayato159/isekai-shop-api/domains/purchasing/controller"
	_purchasingRepository "github.com/Rayato159/isekai-shop-api/domains/purchasing/repository"
	_purchasingService "github.com/Rayato159/isekai-shop-api/domains/purchasing/service"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
)

func (s *echoServer) initPurchasingRouter(customMiddleware customMiddleware.CustomMiddleware) {
	router := s.baseRouter.Group("/balancing")

	balancingRepository := _balancingRepository.NewBalancingRepositoryImpl(s.db, s.app.Logger)
	itemRepository := _itemGettingRepository.NewItemGettingRepositoryImpl(s.db, s.app.Logger)
	purchasingRepository := _purchasingRepository.NewPurchasingRepositoryImpl(s.db, s.app.Logger)
	inventoryRepository := _playerSouce.NewInventoryRepositoryImpl(s.db, s.app.Logger)

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
