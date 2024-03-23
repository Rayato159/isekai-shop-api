package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"

	"github.com/stretchr/testify/mock"
)

type ItemShopRepositoryMock struct {
	mock.Mock
}

func (m *ItemShopRepositoryMock) FindByID(itemID uint64) (*entities.Item, error) {
	args := m.Called(itemID)
	return args.Get(0).(*entities.Item), args.Error(1)
}

func (m *ItemShopRepositoryMock) Listing(itemFilterDto *entities.ItemFilterDto) ([]*entities.Item, error) {
	args := m.Called(itemFilterDto)
	return args.Get(0).([]*entities.Item), args.Error(1)
}

func (m *ItemShopRepositoryMock) FindByIDList(itemIDs []uint64) ([]*entities.Item, error) {
	args := m.Called(itemIDs)
	return args.Get(0).([]*entities.Item), args.Error(1)
}

func (m *ItemShopRepositoryMock) Counting(itemFilterDto *entities.ItemFilterDto) (int64, error) {
	args := m.Called(itemFilterDto)
	return args.Get(0).(int64), args.Error(1)
}

func (m *ItemShopRepositoryMock) PurchaseHistoryRecording(purchasingEntity *entities.PurchaseHistory) (*entities.PurchaseHistory, error) {
	args := m.Called(purchasingEntity)
	return args.Get(0).(*entities.PurchaseHistory), args.Error(1)
}

func (m *ItemShopRepositoryMock) TransactionBegin() {
	m.Called()
}

func (m *ItemShopRepositoryMock) TransactionRollback() {
	m.Called()
}

func (m *ItemShopRepositoryMock) TransactionCommit() error {
	args := m.Called()
	return args.Error(0)
}
