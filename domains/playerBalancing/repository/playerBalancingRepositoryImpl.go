package repository

import (
	_playerBalancingException "github.com/Rayato159/isekai-shop-api/domains/playerBalancing/exception"
	entities "github.com/Rayato159/isekai-shop-api/entities"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type playerBalancingImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewPlayerBalancingRepositoryImpl(db *gorm.DB, logger echo.Logger) PlayerBalancingRepository {
	return &playerBalancingImpl{
		db:     db,
		logger: logger,
	}
}

func (r *playerBalancingImpl) Recording(playerBalancingEntity *entities.PlayerBalancing) (*entities.PlayerBalancing, error) {
	insertedPlayerBalancing := new(entities.PlayerBalancing)

	if err := r.db.Create(playerBalancingEntity).Scan(insertedPlayerBalancing).Error; err != nil {
		r.logger.Error("Failed to insert balancing", err.Error())
		return nil, &_playerBalancingException.PlayerBalanceRecording{}
	}

	return insertedPlayerBalancing, nil
}

func (r *playerBalancingImpl) Showing(playerID string) (*entities.PlayerBalanceShowingDto, error) {
	balanceDto := new(entities.PlayerBalanceShowingDto)

	if err := r.db.Model(
		&entities.PlayerBalancing{},
	).Where(
		"player_id = ?", playerID,
	).Select(
		"player_id, sum(amount) as balance",
	).Group(
		"player_id",
	).Scan(&balanceDto).Error; err != nil {
		r.logger.Error("Failed to calculate player balance", err.Error())
		return nil, &_playerBalancingException.PlayerBalanceShowingException{PlayerID: playerID}
	}

	return balanceDto, nil
}
