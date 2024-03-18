package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"

	"github.com/stretchr/testify/mock"
)

type ItemGettingRepositoryMock struct {
	mock.Mock
}

func (m *ItemGettingRepositoryMock) FindItemByID(itemID uint64) (*entities.Item, error) {
	args := m.Called(itemID)
	return args.Get(0).(*entities.Item), args.Error(1)
}

func (m *ItemGettingRepositoryMock) ItemListing(itemFilterDto *entities.ItemFilterDto) ([]*entities.Item, error) {
	args := m.Called(itemFilterDto)
	return args.Get(0).([]*entities.Item), args.Error(1)
}

func (m *ItemGettingRepositoryMock) FindItemByIDs(itemIDs []uint64) ([]*entities.Item, error) {
	args := m.Called(itemIDs)
	return args.Get(0).([]*entities.Item), args.Error(1)
}

func (m *ItemGettingRepositoryMock) ItemCounting(itemFilterDto *entities.ItemFilterDto) (int64, error) {
	args := m.Called(itemFilterDto)
	return args.Get(0).(int64), args.Error(1)
}
