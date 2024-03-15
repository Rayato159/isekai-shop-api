package repository

import (
	_historyOfPurchasingEntity "github.com/Rayato159/isekai-shop-api/domains/historyOfPurchasing/entity"

	"github.com/stretchr/testify/mock"
)

type HistoryOfPurchasingRepositoryMock struct {
	mock.Mock
}

func (m *HistoryOfPurchasingRepositoryMock) HistoryOfPurchasingRecording(historyOfPurchasingEntity *_historyOfPurchasingEntity.HistoryOfPurchasing) (*_historyOfPurchasingEntity.HistoryOfPurchasing, error) {
	args := m.Called(historyOfPurchasingEntity)
	return args.Get(0).(*_historyOfPurchasingEntity.HistoryOfPurchasing), args.Error(1)
}

func (m *HistoryOfPurchasingRepositoryMock) PlayerHistoryOfPurchasingListing(playerID string) ([]*_historyOfPurchasingEntity.HistoryOfPurchasing, error) {
	args := m.Called(playerID)
	return args.Get(0).([]*_historyOfPurchasingEntity.HistoryOfPurchasing), args.Error(1)
}
