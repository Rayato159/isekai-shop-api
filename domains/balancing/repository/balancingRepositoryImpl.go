package repository

import (
	_balancingException "github.com/Rayato159/isekai-shop-api/domains/balancing/exception"
	entities "github.com/Rayato159/isekai-shop-api/entities"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type balancingRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewBalancingRepositoryImpl(db *gorm.DB, logger echo.Logger) BalancingRepository {
	return &balancingRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *balancingRepositoryImpl) BalancingRecording(balancingEntity *entities.Balancing) (*entities.Balancing, error) {
	insertedBalancing := new(entities.Balancing)

	if err := r.db.Create(balancingEntity).Scan(insertedBalancing).Error; err != nil {
		r.logger.Error("Failed to insert balancing", err.Error())
		return nil, &_balancingException.BalancingRecordingException{}
	}

	return insertedBalancing, nil
}

func (r *balancingRepositoryImpl) PlayerBalanceShowing(playerID string) (*entities.PlayerBalanceDto, error) {
	balanceDto := new(entities.PlayerBalanceDto)

	if err := r.db.Model(
		&entities.Balancing{},
	).Where(
		"player_id = ?", playerID,
	).Select(
		"player_id, sum(amount) as balance",
	).Group(
		"player_id",
	).Scan(&balanceDto).Error; err != nil {
		r.logger.Error("Failed to calculate player balance", err.Error())
		return nil, &_balancingException.PlayerBalanceShowingException{PlayerID: playerID}
	}

	return balanceDto, nil
}
