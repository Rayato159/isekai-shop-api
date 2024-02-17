package server

import (
	_orderController "github.com/Rayato159/isekai-shop-api/modules/order/controller"
	_orderRepository "github.com/Rayato159/isekai-shop-api/modules/order/repository"
	_orderService "github.com/Rayato159/isekai-shop-api/modules/order/service"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
)

func (s *echoServer) initOrderRouter(customMiddleware customMiddleware.CustomMiddleware) {
	router := s.baseRouter.Group("/order")

	orderRepository := _orderRepository.NewOrderRepository(s.db, s.app.Logger)
	orderService := _orderService.NewOrderServiceImpl(orderRepository)
	orderController := _orderController.NewOrderControllerImpl(orderService, s.app.Logger)

	router.GET("", orderController.PlayerOrderListing, customMiddleware.PlayerAuthorize)
}
