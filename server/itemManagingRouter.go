package server

import (
	_itemManagingController "github.com/Rayato159/isekai-shop-api/pkg/itemManaging/controller"
	_itemManagingRepository "github.com/Rayato159/isekai-shop-api/pkg/itemManaging/repository"
	_itemManagingService "github.com/Rayato159/isekai-shop-api/pkg/itemManaging/service"
	_itemShopRepository "github.com/Rayato159/isekai-shop-api/pkg/itemShop/repository"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
)

func (s *echoServer) initItemManagingRouter(customMiddleware customMiddleware.CustomMiddleware) {
	router := s.baseRouter.Group("/item-managing")

	itemRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemMangingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)

	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(itemMangingRepository, itemRepository)

	itemManaging := _itemManagingController.NewItemManagingControllerImpl(itemManagingService)

	router.POST("", itemManaging.Creating, customMiddleware.AdminAuthorizing)
	router.PATCH("/:itemID", itemManaging.Editing, customMiddleware.AdminAuthorizing)
	router.DELETE("/:itemID", itemManaging.Archiving, customMiddleware.AdminAuthorizing)
}
