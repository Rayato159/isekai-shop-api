package server

import (
	_adminController "github.com/Rayato159/isekai-shop-api/modules/admin/controller"
	_adminService "github.com/Rayato159/isekai-shop-api/modules/admin/service"
	_itemRepository "github.com/Rayato159/isekai-shop-api/modules/item/repository"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
)

func (s *echoServer) initAdminRouter(customMiddleware customMiddleware.CustomMiddleware) {
	router := s.baseRouter.Group("/admin/item")

	itemRepository := _itemRepository.NewItemRepositoryImpl(s.db, s.app.Logger)
	adminService := _adminService.NewAdminServiceImpl(itemRepository)
	adminController := _adminController.NewAdminControllerImpl(adminService, s.app.Logger)

	router.POST("", adminController.ItemCreating, customMiddleware.AdminAuthorize)
	router.PATCH("/:itemID", adminController.ItemEditing, customMiddleware.AdminAuthorize)
	router.DELETE("/:itemID", adminController.ItemArchiving, customMiddleware.AdminAuthorize)
}
