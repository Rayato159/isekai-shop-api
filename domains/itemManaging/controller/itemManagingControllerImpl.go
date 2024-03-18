package controller

import (
	"net/http"
	"strconv"

	_itemManagingModel "github.com/Rayato159/isekai-shop-api/domains/itemManaging/model"
	_itemManging "github.com/Rayato159/isekai-shop-api/domains/itemManaging/service"
	"github.com/Rayato159/isekai-shop-api/domains/utils"
	"github.com/Rayato159/isekai-shop-api/server/writter"
	"github.com/labstack/echo/v4"
)

type adminControllerImpl struct {
	itemManging _itemManging.ItemManagingService
	logger      echo.Logger
}

func NewItemManagingControllerImpl(itemManging _itemManging.ItemManagingService, logger echo.Logger) ItemManagingController {
	return &adminControllerImpl{
		itemManging: itemManging,
		logger:      logger,
	}
}

func (c *adminControllerImpl) ItemCreating(pctx echo.Context) error {
	adminID, err := utils.GetAdminID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusUnauthorized, err)
	}

	itemCreatingReq := new(_itemManagingModel.ItemCreatingReq)

	if err := pctx.Bind(itemCreatingReq); err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}
	itemCreatingReq.AdminID = adminID

	item, err := c.itemManging.ItemCreating(itemCreatingReq)
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

	editItemReq := new(_itemManagingModel.ItemEditingReq)

	if err := pctx.Bind(editItemReq); err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}
	editItemReq.AdminID = adminID

	item, err := c.itemManging.ItemEditing(itemID, editItemReq)
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

	err = c.itemManging.ItemArchiving(itemID)
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
