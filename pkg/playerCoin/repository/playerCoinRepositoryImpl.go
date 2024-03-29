package repository

import (
	"github.com/Rayato159/isekai-shop-api/databases"
	"github.com/Rayato159/isekai-shop-api/entities"
	_playerCoin "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/exception"
	_playerCoinModel "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/model"

	"github.com/labstack/echo/v4"
)

type playerCoinRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerCoinRepositoryImpl(db databases.Database, logger echo.Logger) PlayerCoinRepository {
	return &playerCoinRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *playerCoinRepositoryImpl) CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error) {
	playerCoin := new(entities.PlayerCoin)

	if err := r.db.Connect().Create(playerCoinEntity).Scan(playerCoin).Error; err != nil {
		r.logger.Error("Player's balance recording failed:", err.Error())
		return nil, &_playerCoin.CoinAdding{}
	}

	return playerCoin, nil
}

func (r *playerCoinRepositoryImpl) ReverseCoinAdding(playerCoinEntity *entities.PlayerCoin) error {
	if err := r.db.Connect().Delete(playerCoinEntity).Error; err != nil {
		r.logger.Error("Player's balance deleting failed:", err.Error())
		return &_playerCoin.CoinAdding{}
	}

	return nil
}

func (r *playerCoinRepositoryImpl) Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error) {
	playerCoin := new(_playerCoinModel.PlayerCoinShowing)

	if err := r.db.Connect().Model(
		&entities.PlayerCoin{},
	).Where(
		"player_id = ?", playerID,
	).Select(
		"player_id, sum(amount) as coin",
	).Group(
		"player_id",
	).Scan(&playerCoin).Error; err != nil {
		r.logger.Error("Calculating player coin failed:", err.Error())
		return nil, &_playerCoin.PlayerCoinShowing{}
	}

	return playerCoin, nil
}
