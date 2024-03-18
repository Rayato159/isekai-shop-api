package repository

import (
	_purchasingEntity "github.com/Rayato159/isekai-shop-api/domains/purchasing/entity"

	"github.com/stretchr/testify/mock"
)

type PurchasingRepositoryMock struct {
	mock.Mock
}

func (m *PurchasingRepositoryMock) PurchasingHistoryRecording(purchasingEntity *_purchasingEntity.Purchasing) (*_purchasingEntity.Purchasing, error) {
	args := m.Called(purchasingEntity)
	return args.Get(0).(*_purchasingEntity.Purchasing), args.Error(1)
}
