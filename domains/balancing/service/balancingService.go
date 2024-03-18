package service

import (
	_balancingModel "github.com/Rayato159/isekai-shop-api/domains/balancing/model"
)

type BalancingService interface {
	TopUp(topUpReq *_balancingModel.TopUpReq) (*_balancingModel.Balancing, error)
	PlayerBalanceShowing(playerID string) *_balancingModel.PlayerBalance
}
