package repository

import (
	"github.com/Rayato159/isekai-shop-api/databases"
	entities "github.com/Rayato159/isekai-shop-api/entities"
	_itemShopModel "github.com/Rayato159/isekai-shop-api/pkg/itemShop/model"

	"github.com/stretchr/testify/mock"
)

type ItemShopRepositoryMock struct {
	mock.Mock
}

func (m *ItemShopRepositoryMock) FindByID(itemID uint64) (*entities.Item, error) {
	args := m.Called(itemID)
	return args.Get(0).(*entities.Item), args.Error(1)
}

func (m *ItemShopRepositoryMock) Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) {
	args := m.Called(itemFilter)
	return args.Get(0).([]*entities.Item), args.Error(1)
}

func (m *ItemShopRepositoryMock) FindByIDList(itemIDs []uint64) ([]*entities.Item, error) {
	args := m.Called(itemIDs)
	return args.Get(0).([]*entities.Item), args.Error(1)
}

func (m *ItemShopRepositoryMock) Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error) {
	args := m.Called(itemFilter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *ItemShopRepositoryMock) PurchaseHistoryRecording(purchasingEntity *entities.PurchaseHistory) (*entities.PurchaseHistory, error) {
	args := m.Called(purchasingEntity)
	return args.Get(0).(*entities.PurchaseHistory), args.Error(1)
}

func (m *ItemShopRepositoryMock) Transaction() databases.Database {
	args := m.Called()
	return args.Get(0).(databases.Database)
}
