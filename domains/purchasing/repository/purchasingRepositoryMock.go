package repository

import (
	_purchasingEntity "github.com/Rayato159/isekai-shop-api/domains/purchasing/entity"

	"github.com/stretchr/testify/mock"
)

type PurchasingRepositoryMock struct {
	mock.Mock
}

func (m *PurchasingRepositoryMock) PurchasingHistoryRecording(purchasingEntity *_purchasingEntity.PurchasingHistory) (*_purchasingEntity.PurchasingHistory, error) {
	args := m.Called(purchasingEntity)
	return args.Get(0).(*_purchasingEntity.PurchasingHistory), args.Error(1)
}
