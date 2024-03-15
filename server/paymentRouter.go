package server

import (
	_historyOfPurchasingRepository "github.com/Rayato159/isekai-shop-api/domains/historyOfPurchasing/repository"
	_itemRepository "github.com/Rayato159/isekai-shop-api/domains/item/repository"
	_paymentController "github.com/Rayato159/isekai-shop-api/domains/payment/controller"
	_paymentRepository "github.com/Rayato159/isekai-shop-api/domains/payment/repository"
	_paymentService "github.com/Rayato159/isekai-shop-api/domains/payment/service"
	_playerSouce "github.com/Rayato159/isekai-shop-api/domains/player/repository"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
)

func (s *echoServer) initPaymentRouter(customMiddleware customMiddleware.CustomMiddleware) {
	router := s.baseRouter.Group("/payment")

	paymentRepository := _paymentRepository.NewPaymentRepositoryImpl(s.db, s.app.Logger)
	itemRepository := _itemRepository.NewItemRepositoryImpl(s.db, s.app.Logger)
	historyOfPurchasingRepository := _historyOfPurchasingRepository.NewHistoryOfPurchasingRepositoryImpl(s.db, s.app.Logger)
	inventoryRepository := _playerSouce.NewInventoryRepositoryImpl(s.db, s.app.Logger)

	paymentService := _paymentService.NewPaymentServiceImpl(
		paymentRepository,
		itemRepository,
		historyOfPurchasingRepository,
		inventoryRepository,
	)
	paymentController := _paymentController.NewPaymentControllerImpl(paymentService, s.app.Logger)

	router.POST("", paymentController.TopUp, customMiddleware.PlayerAuthorizing)
	router.GET("/balance", paymentController.PlayerBalanceShowing, customMiddleware.PlayerAuthorizing)
	router.POST("/buy", paymentController.ItemBuying, customMiddleware.PlayerAuthorizing)
	router.POST("/sell", paymentController.ItemSelling, customMiddleware.PlayerAuthorizing)
}
