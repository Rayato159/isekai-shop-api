package repository

import (
	_inventoryEntity "github.com/Rayato159/isekai-shop-api/modules/inventory/entity"

	"github.com/stretchr/testify/mock"
)

type InventoryRepositoryMock struct {
	mock.Mock
}

func (m *InventoryRepositoryMock) InsertInventoryInBluk(
	inventoryEntities []*_inventoryEntity.Inventory,
) ([]*_inventoryEntity.Inventory, error) {
	args := m.Called(inventoryEntities)
	return args.Get(0).([]*_inventoryEntity.Inventory), args.Error(1)
}

func (m *InventoryRepositoryMock) FindPlayerInventories(playerID string) ([]*_inventoryEntity.Inventory, error) {
	args := m.Called(playerID)
	return args.Get(0).([]*_inventoryEntity.Inventory), args.Error(1)
}

func (m *InventoryRepositoryMock) DeleteItemByLimit(playerID string, itemID uint64, limit int) error {
	args := m.Called(playerID, itemID, limit)
	return args.Error(0)
}

func (m *InventoryRepositoryMock) CountPlayerItem(playerID string, itemID uint64) int64 {
	args := m.Called(playerID, itemID)
	return args.Get(0).(int64)
}
