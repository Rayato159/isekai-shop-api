package controller

import (
	"net/http"
	"strconv"

	_adminService "github.com/Rayato159/isekai-shop-api/modules/admin/service"
	_itemModel "github.com/Rayato159/isekai-shop-api/modules/item/model"
	"github.com/Rayato159/isekai-shop-api/modules/utils"
	"github.com/Rayato159/isekai-shop-api/server/writter"
	"github.com/labstack/echo/v4"
)

type adminControllerImpl struct {
	adminService _adminService.AdminService
	logger       echo.Logger
}

func NewAdminControllerImpl(adminService _adminService.AdminService, logger echo.Logger) AdminController {
	return &adminControllerImpl{
		adminService: adminService,
		logger:       logger,
	}
}

func (c *adminControllerImpl) ItemCreating(pctx echo.Context) error {
	adminID, err := utils.GetAdminID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusUnauthorized, err)
	}

	createItemReq := new(_itemModel.ItemCreatingReq)

	if err := pctx.Bind(createItemReq); err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}
	createItemReq.AdminID = adminID

	item, err := c.adminService.ItemCreating(createItemReq)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusCreated, item)
}

func (c *adminControllerImpl) ItemEditing(pctx echo.Context) error {
	adminID, err := utils.GetAdminID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusUnauthorized, err)
	}

	itemID, err := c.getItemID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	editItemReq := new(_itemModel.ItemEditingReq)

	if err := pctx.Bind(editItemReq); err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}
	editItemReq.AdminID = adminID

	item, err := c.adminService.ItemEditing(itemID, editItemReq)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, item)
}

func (c *adminControllerImpl) ItemArchiving(pctx echo.Context) error {
	_, err := utils.GetAdminID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusUnauthorized, err)
	}

	itemID, err := c.getItemID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	err = c.adminService.ItemArchiving(itemID)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.NoContent(http.StatusNoContent)
}

func (c *adminControllerImpl) getItemID(pctx echo.Context) (uint64, error) {
	itemID := pctx.Param("itemID")
	itemIDUint64, err := strconv.ParseUint(itemID, 10, 64)
	if err != nil {
		return 0, err
	}

	return itemIDUint64, nil
}
