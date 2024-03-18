package repository

import (
	_playerCoinException "github.com/Rayato159/isekai-shop-api/domains/playerCoin/exception"
	entities "github.com/Rayato159/isekai-shop-api/entities"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type playerCoinImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewPlayerCoinRepositoryImpl(db *gorm.DB, logger echo.Logger) PlayerCoinRepository {
	return &playerCoinImpl{
		db:     db,
		logger: logger,
	}
}

func (r *playerCoinImpl) Recording(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error) {
	insertedPlayerCoin := new(entities.PlayerCoin)

	if err := r.db.Create(playerCoinEntity).Scan(insertedPlayerCoin).Error; err != nil {
		r.logger.Error("Failed to insert balancing", err.Error())
		return nil, &_playerCoinException.PlayerBalanceRecording{}
	}

	return insertedPlayerCoin, nil
}

func (r *playerCoinImpl) Showing(playerID string) (*entities.PlayerCoinShowingDto, error) {
	coinDto := new(entities.PlayerCoinShowingDto)

	if err := r.db.Model(
		&entities.PlayerCoin{},
	).Where(
		"player_id = ?", playerID,
	).Select(
		"player_id, sum(amount) as coin",
	).Group(
		"player_id",
	).Scan(&coinDto).Error; err != nil {
		r.logger.Error("Failed to calculate player coin", err.Error())
		return nil, &_playerCoinException.PlayerCoinShowingException{PlayerID: playerID}
	}

	return coinDto, nil
}
