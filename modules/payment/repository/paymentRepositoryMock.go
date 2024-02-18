package repository

import (
	_paymentEntity "github.com/Rayato159/isekai-shop-api/modules/payment/entity"

	"github.com/stretchr/testify/mock"
)

type PaymentRepositoryMock struct {
	mock.Mock
}

func (m *PaymentRepositoryMock) InsertPayment(paymentEntity *_paymentEntity.Payment) (*_paymentEntity.Payment, error) {
	args := m.Called(paymentEntity)
	return args.Get(0).(*_paymentEntity.Payment), args.Error(1)
}

func (m *PaymentRepositoryMock) CalculatePlayerBalance(playerID string) (*_paymentEntity.PlayerBalanceDto, error) {
	args := m.Called(playerID)
	return args.Get(0).(*_paymentEntity.PlayerBalanceDto), args.Error(1)
}
