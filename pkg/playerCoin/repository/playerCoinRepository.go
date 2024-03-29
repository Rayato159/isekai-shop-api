package repository

import (
	"github.com/Rayato159/isekai-shop-api/entities"
	_playerCoinModel "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/model"
)

type PlayerCoinRepository interface {
	CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
	ReverseCoinAdding(playerCoinEntity *entities.PlayerCoin) error
	Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error)
}
