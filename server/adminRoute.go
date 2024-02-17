package server

import (
	_adminController "github.com/Rayato159/isekai-shop-api/modules/admin/controller"
	_adminService "github.com/Rayato159/isekai-shop-api/modules/admin/service"
	_itemRepository "github.com/Rayato159/isekai-shop-api/modules/item/repository"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
)

func (s *echoServer) initAdminRoute(customMiddleware customMiddleware.CustomMiddleware) {
	router := s.baseRouter.Group("/admin/item")

	itemRepository := _itemRepository.NewItemRepositoryImpl(s.db, s.app.Logger)
	adminService := _adminService.NewAdminServiceImpl(itemRepository, s.app.Logger)
	adminController := _adminController.NewAdminControllerImpl(adminService, s.app.Logger)

	router.POST("", adminController.CreateItem, customMiddleware.AdminAuthorize)
	router.PATCH("/:itemID", adminController.EditItem, customMiddleware.AdminAuthorize)
	router.DELETE("/:itemID", adminController.ArchiveItem, customMiddleware.AdminAuthorize)
}
