package service

import (
	_paymentModel "github.com/Rayato159/isekai-shop-api/modules/payment/model"
)

type PaymentService interface {
	TopUp(topUpReq *_paymentModel.TopUpReq) (*_paymentModel.Payment, error)
	CalculatePlayerBalance(playerID string) *_paymentModel.PlayerBalance
}
