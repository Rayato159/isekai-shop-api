package service

import (
	_orderModel "github.com/Rayato159/isekai-shop-api/modules/order/model"
	_orderRepository "github.com/Rayato159/isekai-shop-api/modules/order/repository"
)

type orderServiceImpl struct {
	orderRepository _orderRepository.OrderRepository
}

func NewOrderServiceImpl(orderRepository _orderRepository.OrderRepository) OrderService {
	return &orderServiceImpl{
		orderRepository: orderRepository,
	}
}

func (s *orderServiceImpl) PlayerOrderListing(playerID string) ([]*_orderModel.Order, error) {
	orderEntites, err := s.orderRepository.FindPlayerOrders(playerID)
	if err != nil {
		return nil, err
	}

	orderModels := make([]*_orderModel.Order, 0)
	for _, order := range orderEntites {
		orderModels = append(orderModels, order.ToOrderModel())
	}

	return orderModels, nil
}
