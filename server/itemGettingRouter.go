package server

import (
	_itemGettingController "github.com/Rayato159/isekai-shop-api/domains/itemGetting/controller"
	_itemGettingRepository "github.com/Rayato159/isekai-shop-api/domains/itemGetting/repository"
	_itemGettingService "github.com/Rayato159/isekai-shop-api/domains/itemGetting/service"
)

func (s *echoServer) initItemRouter() {
	router := s.baseRouter.Group("/item-getting")

	itemRepository := _itemGettingRepository.NewItemGettingRepositoryImpl(s.db, s.app.Logger)
	itemService := _itemGettingService.NewItemServiceImpl(itemRepository)
	itemController := _itemGettingController.NewItemControllerImpl(itemService, s.app.Logger)

	router.GET("", itemController.ItemListing)
}
