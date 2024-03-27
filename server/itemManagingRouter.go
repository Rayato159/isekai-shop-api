package server

import (
	_itemManagingController "github.com/Rayato159/isekai-shop-api/pkg/itemManaging/controller"
	_itemManagingRepository "github.com/Rayato159/isekai-shop-api/pkg/itemManaging/repository"
	_itemManagingService "github.com/Rayato159/isekai-shop-api/pkg/itemManaging/service"
	_itemShopRepository "github.com/Rayato159/isekai-shop-api/pkg/itemShop/repository"
)

func (s *echoServer) initItemManagingRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/item-managing")

	itemRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemMangingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)

	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(itemMangingRepository, itemRepository)

	itemManaging := _itemManagingController.NewItemManagingControllerImpl(itemManagingService)

	router.POST("", itemManaging.Creating, m.AdminAuthorizing)
	router.PATCH("/:itemID", itemManaging.Editing, m.AdminAuthorizing)
	router.DELETE("/:itemID", itemManaging.Archiving, m.AdminAuthorizing)
}
