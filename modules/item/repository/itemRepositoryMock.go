package repository

import (
	_itemEntity "github.com/Rayato159/isekai-shop-api/modules/item/entity"

	"github.com/stretchr/testify/mock"
)

type ItemRepositoryMock struct {
	mock.Mock
}

func (m *ItemRepositoryMock) FindItemByID(itemID string) (*_itemEntity.Item, error) {
	args := m.Called(itemID)
	return args.Get(0).(*_itemEntity.Item), args.Error(1)
}

func (m *ItemRepositoryMock) FindItems(itemFilterDto *_itemEntity.ItemFilterDto) ([]*_itemEntity.Item, error) {
	args := m.Called(itemFilterDto)
	return args.Get(0).([]*_itemEntity.Item), args.Error(1)
}

func (m *ItemRepositoryMock) FindItemByIDs(itemIDs []string) ([]*_itemEntity.Item, error) {
	args := m.Called(itemIDs)
	return args.Get(0).([]*_itemEntity.Item), args.Error(1)
}

func (m *ItemRepositoryMock) CountItems(itemFilterDto *_itemEntity.ItemFilterDto) (int64, error) {
	args := m.Called(itemFilterDto)
	return args.Get(0).(int64), args.Error(1)
}

func (m *ItemRepositoryMock) InsertItem(itemEntity *_itemEntity.Item) (*_itemEntity.Item, error) {
	args := m.Called(itemEntity)
	return args.Get(0).(*_itemEntity.Item), args.Error(1)
}

func (m *ItemRepositoryMock) UpdateItem(itemID string, updateItemDto *_itemEntity.UpdateItemDto) (string, error) {
	args := m.Called(itemID, updateItemDto)
	return args.Get(0).(string), args.Error(1)
}

func (m *ItemRepositoryMock) ArchiveItem(itemID string) error {
	args := m.Called(itemID)
	return args.Error(0)
}
