package server

import (
	_paymentController "github.com/Rayato159/isekai-shop-api/modules/payment/controller"
	_paymentRepository "github.com/Rayato159/isekai-shop-api/modules/payment/repository"
	_paymentService "github.com/Rayato159/isekai-shop-api/modules/payment/service"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
)

func (s *echoServer) initPaymentRouter(customMiddleware customMiddleware.CustomMiddleware) {
	router := s.baseRouter.Group("/payment")

	paymentRepository := _paymentRepository.NewPaymentRepositoryImpl(s.db, s.app.Logger)
	paymentService := _paymentService.NewPaymentServiceImpl(paymentRepository, s.app.Logger)
	paymentController := _paymentController.NewPaymentControllerImpl(paymentService, s.app.Logger)

	router.POST("", paymentController.TopUp, customMiddleware.PlayerAuthorize)
	router.GET("/balance", paymentController.CalculatePlayerBalance, customMiddleware.PlayerAuthorize)
}
