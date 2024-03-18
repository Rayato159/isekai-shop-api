package server

import (
	_itemManagingController "github.com/Rayato159/isekai-shop-api/domains/itemManaging/controller"
	_itemManagingRepository "github.com/Rayato159/isekai-shop-api/domains/itemManaging/repository"
	_itemManagingService "github.com/Rayato159/isekai-shop-api/domains/itemManaging/service"
	_itemShopRepository "github.com/Rayato159/isekai-shop-api/domains/itemShop/repository"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
)

func (s *echoServer) initItemManagingRouter(customMiddleware customMiddleware.CustomMiddleware) {
	router := s.baseRouter.Group("/item-managing")

	itemRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemMangingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)

	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(itemMangingRepository, itemRepository)

	itemManaging := _itemManagingController.NewItemManagingControllerImpl(itemManagingService, s.app.Logger)

	router.POST("", itemManaging.ItemCreating, customMiddleware.AdminAuthorizing)
	router.PATCH("/:itemID", itemManaging.ItemEditing, customMiddleware.AdminAuthorizing)
	router.DELETE("/:itemID", itemManaging.ItemArchiving, customMiddleware.AdminAuthorizing)
}
