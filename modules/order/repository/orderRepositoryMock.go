package repository

import (
	_orderEntity "github.com/Rayato159/isekai-shop-api/modules/order/entity"

	"github.com/stretchr/testify/mock"
)

type OrderRepositoryMock struct {
	mock.Mock
}

func (m *OrderRepositoryMock) InsertOrder(orderEntity *_orderEntity.Order) (*_orderEntity.Order, error) {
	args := m.Called(orderEntity)
	return args.Get(0).(*_orderEntity.Order), args.Error(1)
}

func (m *OrderRepositoryMock) FindPlayerOrders(playerID string) ([]*_orderEntity.Order, error) {
	args := m.Called(playerID)
	return args.Get(0).([]*_orderEntity.Order), args.Error(1)
}
