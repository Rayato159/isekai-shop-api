package service

import (
	_playerCoinModel "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/model"
)

type PlayerCoinService interface {
	CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin, error)
	PlayerCoinShowing(playerID string) *_playerCoinModel.PlayerCoinShowing
}
