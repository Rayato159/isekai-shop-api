package repository

import (
	_orderEntity "github.com/Rayato159/isekai-shop-api/modules/order/entity"
)

type OrderRepository interface {
	InsertOrder(orderEntity *_orderEntity.Order) (*_orderEntity.Order, error)
	FindPlayerOrders(playerID string) ([]*_orderEntity.Order, error)
}
