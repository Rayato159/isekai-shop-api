package controller

import (
	"net/http"
	"strconv"

	"github.com/Rayato159/isekai-shop-api/domains/common"
	_itemManagingModel "github.com/Rayato159/isekai-shop-api/domains/itemManaging/model"
	_itemManging "github.com/Rayato159/isekai-shop-api/domains/itemManaging/service"
	"github.com/Rayato159/isekai-shop-api/server/validation"
	"github.com/Rayato159/isekai-shop-api/server/writter"
	"github.com/labstack/echo/v4"
)

type itemManagingImpl struct {
	itemManging _itemManging.ItemManagingService
}

func NewItemManagingControllerImpl(itemManging _itemManging.ItemManagingService) ItemManagingController {
	return &itemManagingImpl{itemManging: itemManging}
}

func (c *itemManagingImpl) Creating(pctx echo.Context) error {
	adminID, err := common.GetAdminID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusUnauthorized, err)
	}

	itemCreatingReq := new(_itemManagingModel.ItemCreatingReq)

	validatingContext := validation.NewCustomEchoRequest(pctx)

	if err := validatingContext.Bind(itemCreatingReq); err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}
	itemCreatingReq.AdminID = adminID

	item, err := c.itemManging.Creating(itemCreatingReq)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusCreated, item)
}

func (c *itemManagingImpl) Editing(pctx echo.Context) error {
	adminID, err := common.GetAdminID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusUnauthorized, err)
	}

	itemID, err := c.getItemID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	editItemReq := new(_itemManagingModel.ItemEditingReq)

	validatingContext := validation.NewCustomEchoRequest(pctx)

	if err := validatingContext.Bind(editItemReq); err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}
	editItemReq.AdminID = adminID

	item, err := c.itemManging.Editing(itemID, editItemReq)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, item)
}

func (c *itemManagingImpl) Archiving(pctx echo.Context) error {
	_, err := common.GetAdminID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusUnauthorized, err)
	}

	itemID, err := c.getItemID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	err = c.itemManging.Archiving(itemID)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.NoContent(http.StatusNoContent)
}

func (c *itemManagingImpl) getItemID(pctx echo.Context) (uint64, error) {
	itemID := pctx.Param("itemID")
	itemIDUint64, err := strconv.ParseUint(itemID, 10, 64)
	if err != nil {
		return 0, err
	}

	return itemIDUint64, nil
}
