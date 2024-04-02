package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"
	"gorm.io/gorm"

	"github.com/stretchr/testify/mock"
)

type InventoryRepositoryMock struct {
	mock.Mock
}

func (m *InventoryRepositoryMock) Filling(playerID string, itemID uint64, qty int, tx *gorm.DB) ([]*entities.Inventory, error) {
	args := m.Called(playerID, itemID, qty, tx)
	return args.Get(0).([]*entities.Inventory), args.Error(1)
}

func (m *InventoryRepositoryMock) Listing(playerID string) ([]*entities.Inventory, error) {
	args := m.Called(playerID)
	return args.Get(0).([]*entities.Inventory), args.Error(1)
}

func (m *InventoryRepositoryMock) Removing(playerID string, itemID uint64, limit int, tx *gorm.DB) error {
	args := m.Called(playerID, itemID, limit, tx)
	return args.Error(0)
}

func (m *InventoryRepositoryMock) PlayerItemCounting(playerID string, itemID uint64) int64 {
	args := m.Called(playerID, itemID)
	return args.Get(0).(int64)
}
