package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"

	"github.com/stretchr/testify/mock"
)

type ItemManagingRepositoryMock struct {
	mock.Mock
}

func (m *ItemManagingRepositoryMock) ItemCreating(itemEntity *entities.Item) (*entities.Item, error) {
	args := m.Called(itemEntity)
	return args.Get(0).(*entities.Item), args.Error(1)
}

func (m *ItemManagingRepositoryMock) ItemEditing(itemID uint64, updateItemDto *entities.ItemEditingDto) (uint64, error) {
	args := m.Called(itemID, updateItemDto)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *ItemManagingRepositoryMock) ItemArchiving(itemID uint64) error {
	args := m.Called(itemID)
	return args.Error(0)
}
