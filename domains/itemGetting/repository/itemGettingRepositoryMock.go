package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"

	"github.com/stretchr/testify/mock"
)

type ItemGettingRepositoryMock struct {
	mock.Mock
}

func (m *ItemGettingRepositoryMock) FindByID(itemID uint64) (*entities.Item, error) {
	args := m.Called(itemID)
	return args.Get(0).(*entities.Item), args.Error(1)
}

func (m *ItemGettingRepositoryMock) Listing(itemFilterDto *entities.ItemFilterDto) ([]*entities.Item, error) {
	args := m.Called(itemFilterDto)
	return args.Get(0).([]*entities.Item), args.Error(1)
}

func (m *ItemGettingRepositoryMock) FindByIDList(itemIDs []uint64) ([]*entities.Item, error) {
	args := m.Called(itemIDs)
	return args.Get(0).([]*entities.Item), args.Error(1)
}

func (m *ItemGettingRepositoryMock) Counting(itemFilterDto *entities.ItemFilterDto) (int64, error) {
	args := m.Called(itemFilterDto)
	return args.Get(0).(int64), args.Error(1)
}
