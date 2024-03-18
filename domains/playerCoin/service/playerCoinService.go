package service

import (
	_playerBalacingModel "github.com/Rayato159/isekai-shop-api/domains/playerCoin/model"
)

type PlayerCoinService interface {
	BuyingCoin(buyingCoinReq *_playerBalacingModel.BuyingCoinReq) (*_playerBalacingModel.PlayerCoin, error)
	PlayerCoinShowing(playerID string) *_playerBalacingModel.PlayerCoinShowing
}
