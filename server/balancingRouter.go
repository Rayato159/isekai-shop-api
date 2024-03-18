package server

import (
	_balancingController "github.com/Rayato159/isekai-shop-api/domains/balancing/controller"
	_balancingRepository "github.com/Rayato159/isekai-shop-api/domains/balancing/repository"
	_balancingService "github.com/Rayato159/isekai-shop-api/domains/balancing/service"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
)

func (s *echoServer) initBalancingRouter(customMiddleware customMiddleware.CustomMiddleware) {
	router := s.baseRouter.Group("/balancing")

	balancingRepository := _balancingRepository.NewBalancingRepositoryImpl(s.db, s.app.Logger)

	balancingService := _balancingService.NewBalancingServiceImpl(
		balancingRepository,
	)
	balancingController := _balancingController.NewBalancingControllerImpl(balancingService, s.app.Logger)

	router.POST("", balancingController.TopUp, customMiddleware.PlayerAuthorizing)
	router.GET("/balance", balancingController.PlayerBalanceShowing, customMiddleware.PlayerAuthorizing)
}
