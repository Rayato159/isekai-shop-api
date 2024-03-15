package controller

import (
	"net/http"

	_paymentModel "github.com/Rayato159/isekai-shop-api/domains/payment/model"
	_paymentService "github.com/Rayato159/isekai-shop-api/domains/payment/service"
	"github.com/Rayato159/isekai-shop-api/domains/utils"
	"github.com/Rayato159/isekai-shop-api/server/writter"
	"github.com/labstack/echo/v4"
)

type paymentControllerImpl struct {
	paymentService _paymentService.PaymentService
	logger         echo.Logger
}

func NewPaymentControllerImpl(paymentService _paymentService.PaymentService, logger echo.Logger) PaymentController {
	return &paymentControllerImpl{
		paymentService: paymentService,
		logger:         logger,
	}
}

func (c *paymentControllerImpl) TopUp(pctx echo.Context) error {
	playerID, err := utils.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	topUpReq := new(_paymentModel.TopUpReq)

	if err := pctx.Bind(topUpReq); err != nil {
		c.logger.Error("Failed to bind top up request", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}
	topUpReq.PlayerID = playerID

	payment, err := c.paymentService.TopUp(topUpReq)
	if err != nil {
		c.logger.Error("Failed to top up", err.Error())
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusCreated, payment)
}

func (c *paymentControllerImpl) PlayerBalanceShowing(pctx echo.Context) error {
	playerID, err := utils.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	balance := c.paymentService.PlayerBalanceShowing(playerID)

	return pctx.JSON(http.StatusOK, balance)
}

func (c *paymentControllerImpl) ItemBuying(pctx echo.Context) error {
	playerID, err := utils.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	itemBuyingReq := new(_paymentModel.ItemBuyingReq)

	if err := pctx.Bind(itemBuyingReq); err != nil {
		c.logger.Error("Failed to bind buy item request", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}
	itemBuyingReq.PlayerID = playerID

	payment, err := c.paymentService.ItemBuying(itemBuyingReq)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, payment)
}

func (c *paymentControllerImpl) ItemSelling(pctx echo.Context) error {
	playerID, err := utils.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	itemSellingReq := new(_paymentModel.ItemSellingReq)

	if err := pctx.Bind(itemSellingReq); err != nil {
		c.logger.Error("Failed to bind sell item request", err.Error())
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}
	itemSellingReq.PlayerID = playerID

	payment, err := c.paymentService.ItemSelling(itemSellingReq)
	if err != nil {
		return writter.CustomError(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, payment)
}
