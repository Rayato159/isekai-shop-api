package repository

import (
	_playerEntity "github.com/Rayato159/isekai-shop-api/modules/player/entity"

	"github.com/stretchr/testify/mock"
)

type InventoryRepositoryMock struct {
	mock.Mock
}

func (m *InventoryRepositoryMock) InventoryFilling(
	inventoryEntities []*_playerEntity.Inventory,
) ([]*_playerEntity.Inventory, error) {
	args := m.Called(inventoryEntities)
	return args.Get(0).([]*_playerEntity.Inventory), args.Error(1)
}

func (m *InventoryRepositoryMock) InventorySearching(playerID string) ([]*_playerEntity.Inventory, error) {
	args := m.Called(playerID)
	return args.Get(0).([]*_playerEntity.Inventory), args.Error(1)
}

func (m *InventoryRepositoryMock) DeleteItemByLimit(playerID string, itemID uint64, limit int) error {
	args := m.Called(playerID, itemID, limit)
	return args.Error(0)
}

func (m *InventoryRepositoryMock) PlayerItemCounting(playerID string, itemID uint64) int64 {
	args := m.Called(playerID, itemID)
	return args.Get(0).(int64)
}
