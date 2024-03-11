package service

import (
	_paymentModel "github.com/Rayato159/isekai-shop-api/modules/payment/model"
)

type PaymentService interface {
	TopUp(topUpReq *_paymentModel.TopUpReq) (*_paymentModel.Payment, error)
	PlayerBalanceShowing(playerID string) *_paymentModel.PlayerBalance
	ItemBuying(itemBuyingReq *_paymentModel.ItemBuyingReq) (*_paymentModel.Payment, error)
	ItemSelling(itemSellingReq *_paymentModel.ItemSellingReq) (*_paymentModel.Payment, error)
}
