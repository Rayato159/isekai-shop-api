package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"

	"github.com/stretchr/testify/mock"
)

type PurchasingRepositoryMock struct {
	mock.Mock
}

func (m *PurchasingRepositoryMock) PurchasingHistoryRecording(purchasingEntity *entities.PurchasingHistory) (*entities.PurchasingHistory, error) {
	args := m.Called(purchasingEntity)
	return args.Get(0).(*entities.PurchasingHistory), args.Error(1)
}
