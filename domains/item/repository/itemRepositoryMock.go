package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"

	"github.com/stretchr/testify/mock"
)

type ItemRepositoryMock struct {
	mock.Mock
}

func (m *ItemRepositoryMock) FindItemByID(itemID uint64) (*entities.Item, error) {
	args := m.Called(itemID)
	return args.Get(0).(*entities.Item), args.Error(1)
}

func (m *ItemRepositoryMock) ItemListing(itemFilterDto *entities.ItemFilterDto) ([]*entities.Item, error) {
	args := m.Called(itemFilterDto)
	return args.Get(0).([]*entities.Item), args.Error(1)
}

func (m *ItemRepositoryMock) FindItemByIDs(itemIDs []uint64) ([]*entities.Item, error) {
	args := m.Called(itemIDs)
	return args.Get(0).([]*entities.Item), args.Error(1)
}

func (m *ItemRepositoryMock) ItemCounting(itemFilterDto *entities.ItemFilterDto) (int64, error) {
	args := m.Called(itemFilterDto)
	return args.Get(0).(int64), args.Error(1)
}
