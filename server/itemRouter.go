package server

import (
	_itemController "github.com/Rayato159/isekai-shop-api/modules/item/controller"
	_itemRepository "github.com/Rayato159/isekai-shop-api/modules/item/repository"
	_itemService "github.com/Rayato159/isekai-shop-api/modules/item/service"
)

func (s *echoServer) initItemRouter() {
	router := s.baseRouter.Group("/item")

	itemRepository := _itemRepository.NewItemRepositoryImpl(s.db, s.app.Logger)
	itemService := _itemService.NewItemServiceImpl(itemRepository, s.app.Logger)
	itemController := _itemController.NewItemControllerImpl(itemService, s.app.Logger)

	router.GET("", itemController.ItemListing)
}
