package service

import (
	_playerBalacingModel "github.com/Rayato159/isekai-shop-api/domains/playerBalancing/model"
)

type PlayerBalancingService interface {
	TopUp(topUpReq *_playerBalacingModel.TopUpReq) (*_playerBalacingModel.PlayerBalancing, error)
	PlayerBalanceShowing(playerID string) *_playerBalacingModel.PlayerBalanceShowing
}
