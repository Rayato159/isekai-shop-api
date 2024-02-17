package controller

import (
	"net/http"

	_orderSergvice "github.com/Rayato159/isekai-shop-api/modules/order/service"
	"github.com/Rayato159/isekai-shop-api/modules/utils"
	"github.com/Rayato159/isekai-shop-api/server/writter"

	"github.com/labstack/echo/v4"
)

type orderControllerImpl struct {
	orderService _orderSergvice.OrderService
	logger       echo.Logger
}

func NewOrderControllerImpl(orderService _orderSergvice.OrderService, logger echo.Logger) OrderController {
	return &orderControllerImpl{
		orderService: orderService,
		logger:       logger,
	}
}

func (c *orderControllerImpl) PlayerOrderListing(pctx echo.Context) error {
	playerID, err := utils.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	orders, err := c.orderService.PlayerOrderListing(playerID)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, orders)
}
