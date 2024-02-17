package service

import (
	_orderModel "github.com/Rayato159/isekai-shop-api/modules/order/model"
)

type OrderService interface {
	PlayerOrderListing(playerID string) ([]*_orderModel.Order, error)
}
