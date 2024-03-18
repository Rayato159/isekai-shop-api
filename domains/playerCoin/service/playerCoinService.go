package service

import (
	_playerCoinModel "github.com/Rayato159/isekai-shop-api/domains/playerCoin/model"
)

type PlayerCoinService interface {
	BuyingCoin(buyingCoinReq *_playerCoinModel.BuyingCoinReq) (*_playerCoinModel.PlayerCoin, error)
	PlayerCoinShowing(playerID string) *_playerCoinModel.PlayerCoinShowing
}
