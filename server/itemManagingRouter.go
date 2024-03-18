package server

import (
	_itemRepository "github.com/Rayato159/isekai-shop-api/domains/item/repository"
	_itemManagingController "github.com/Rayato159/isekai-shop-api/domains/itemManaging/controller"
	_itemManagingRepository "github.com/Rayato159/isekai-shop-api/domains/itemManaging/repository"
	_itemManagingService "github.com/Rayato159/isekai-shop-api/domains/itemManaging/service"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
)

func (s *echoServer) initItemManagingRouter(customMiddleware customMiddleware.CustomMiddleware) {
	router := s.baseRouter.Group("/admin/item")

	itemRepository := _itemRepository.NewItemRepositoryImpl(s.db, s.app.Logger)
	itemMangingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)

	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(itemMangingRepository, itemRepository)

	itemManaging := _itemManagingController.NewItemManagingControllerImpl(itemManagingService, s.app.Logger)

	router.POST("", itemManaging.ItemCreating, customMiddleware.AdminAuthorizing)
	router.PATCH("/:itemID", itemManaging.ItemEditing, customMiddleware.AdminAuthorizing)
	router.DELETE("/:itemID", itemManaging.ItemArchiving, customMiddleware.AdminAuthorizing)
}
