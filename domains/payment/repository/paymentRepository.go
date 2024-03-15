package repository

import _paymentEntity "github.com/Rayato159/isekai-shop-api/domains/payment/entity"

type PaymentRepository interface {
	PaymentRecording(paymentEntity *_paymentEntity.Payment) (*_paymentEntity.Payment, error)
	PlayerBalanceShowing(playerID string) (*_paymentEntity.PlayerBalanceDto, error)
}
