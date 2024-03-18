package repository

import (
	_balancingEntity "github.com/Rayato159/isekai-shop-api/domains/balancing/entity"
	_balancingException "github.com/Rayato159/isekai-shop-api/domains/balancing/exception"

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

func (r *balancingRepositoryImpl) BalancingRecording(balancingEntity *_balancingEntity.Balancing) (*_balancingEntity.Balancing, error) {
	insertedBalancing := new(_balancingEntity.Balancing)

	if err := r.db.Create(balancingEntity).Scan(insertedBalancing).Error; err != nil {
		r.logger.Error("Failed to insert balancing", err.Error())
		return nil, &_balancingException.BalancingRecordingException{}
	}

	return insertedBalancing, nil
}

func (r *balancingRepositoryImpl) PlayerBalanceShowing(playerID string) (*_balancingEntity.PlayerBalanceDto, error) {
	balanceDto := new(_balancingEntity.PlayerBalanceDto)

	if err := r.db.Model(
		&_balancingEntity.Balancing{},
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
