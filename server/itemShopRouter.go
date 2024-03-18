package server

import (
	_itemShopController "github.com/Rayato159/isekai-shop-api/domains/itemShop/controller"
	_itemShopRepository "github.com/Rayato159/isekai-shop-api/domains/itemShop/repository"
	_itemShopService "github.com/Rayato159/isekai-shop-api/domains/itemShop/service"
)

func (s *echoServer) initItemShopRouter() {
	router := s.baseRouter.Group("/item-shop")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemShopService := _itemShopService.NewItemServiceImpl(itemShopRepository)
	itemShopController := _itemShopController.NewItemControllerImpl(itemShopService, s.app.Logger)

	router.GET("", itemShopController.Listing)
}
