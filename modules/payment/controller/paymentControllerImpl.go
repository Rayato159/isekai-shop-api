package controller

import (
	"net/http"

	_paymentModel "github.com/Rayato159/isekai-shop-api/modules/payment/model"
	_paymentService "github.com/Rayato159/isekai-shop-api/modules/payment/service"
	"github.com/Rayato159/isekai-shop-api/modules/utils"
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

func (c *paymentControllerImpl) CalculatePlayerBalance(pctx echo.Context) error {
	playerID, err := utils.GetPlayerID(pctx)
	if err != nil {
		return writter.CustomError(pctx, http.StatusBadRequest, err)
	}

	balance := c.paymentService.CalculatePlayerBalance(playerID)

	return pctx.JSON(http.StatusOK, balance)
}