package server

import (
	_balancingController "github.com/Rayato159/isekai-shop-api/domains/balancing/controller"
	_balancingRepository "github.com/Rayato159/isekai-shop-api/domains/balancing/repository"
	_balancingService "github.com/Rayato159/isekai-shop-api/domains/balancing/service"
	_itemRepository "github.com/Rayato159/isekai-shop-api/domains/item/repository"
	_playerSouce "github.com/Rayato159/isekai-shop-api/domains/player/repository"
	_purchasingRepository "github.com/Rayato159/isekai-shop-api/domains/purchasing/repository"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
)

func (s *echoServer) initBalancingRouter(customMiddleware customMiddleware.CustomMiddleware) {
	router := s.baseRouter.Group("/balancing")

	balancingRepository := _balancingRepository.NewBalancingRepositoryImpl(s.db, s.app.Logger)
	itemRepository := _itemRepository.NewItemRepositoryImpl(s.db, s.app.Logger)
	purchasingRepository := _purchasingRepository.NewPurchasingRepositoryImpl(s.db, s.app.Logger)
	inventoryRepository := _playerSouce.NewInventoryRepositoryImpl(s.db, s.app.Logger)

	balancingService := _balancingService.NewBalancingServiceImpl(
		balancingRepository,
		itemRepository,
		purchasingRepository,
		inventoryRepository,
	)
	balancingController := _balancingController.NewBalancingControllerImpl(balancingService, s.app.Logger)

	router.POST("", balancingController.TopUp, customMiddleware.PlayerAuthorizing)
	router.GET("/balance", balancingController.PlayerBalanceShowing, customMiddleware.PlayerAuthorizing)
	router.POST("/buy", balancingController.ItemBuying, customMiddleware.PlayerAuthorizing)
	router.POST("/sell", balancingController.ItemSelling, customMiddleware.PlayerAuthorizing)
}
