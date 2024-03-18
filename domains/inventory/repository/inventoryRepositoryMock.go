package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"

	"github.com/stretchr/testify/mock"
)

type InventoryRepositoryMock struct {
	mock.Mock
}

func (m *InventoryRepositoryMock) Filling(
	inventoryEntities []*entities.Inventory,
) ([]*entities.Inventory, error) {
	args := m.Called(inventoryEntities)
	return args.Get(0).([]*entities.Inventory), args.Error(1)
}

func (m *InventoryRepositoryMock) Listing(playerID string) ([]*entities.Inventory, error) {
	args := m.Called(playerID)
	return args.Get(0).([]*entities.Inventory), args.Error(1)
}

func (m *InventoryRepositoryMock) Removing(playerID string, itemID uint64, limit int) error {
	args := m.Called(playerID, itemID, limit)
	return args.Error(0)
}

func (m *InventoryRepositoryMock) PlayerItemCounting(playerID string, itemID uint64) int64 {
	args := m.Called(playerID, itemID)
	return args.Get(0).(int64)
}
