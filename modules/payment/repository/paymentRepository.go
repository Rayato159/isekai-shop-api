package repository

import _paymentEntity "github.com/Rayato159/isekai-shop-api/modules/payment/entity"

type PaymentRepository interface {
	TopUp(paymentEntity *_paymentEntity.Payment) (*_paymentEntity.Payment, error)
	CalculatePlayerBalance(playerID string) (*_paymentEntity.PlayerBalanceDto, error)
}
