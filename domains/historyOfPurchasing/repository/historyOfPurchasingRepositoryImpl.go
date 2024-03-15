package repository

import (
	_historyOfPurchasingEntity "github.com/Rayato159/isekai-shop-api/domains/historyOfPurchasing/entity"
	_historyOfPurchasingException "github.com/Rayato159/isekai-shop-api/domains/historyOfPurchasing/exception"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type historyOfPurchasingRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewHistoryOfPurchasingRepositoryImpl(db *gorm.DB, logger echo.Logger) HistoryOfPurchasingRepository {
	return &historyOfPurchasingRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *historyOfPurchasingRepositoryImpl) HistoryOfPurchasingRecording(historyOfPurchasingEntity *_historyOfPurchasingEntity.HistoryOfPurchasing) (*_historyOfPurchasingEntity.HistoryOfPurchasing, error) {
	insertedHistoryOfPurchasing := new(_historyOfPurchasingEntity.HistoryOfPurchasing)

	if err := r.db.Create(historyOfPurchasingEntity).Scan(insertedHistoryOfPurchasing).Error; err != nil {
		r.logger.Errorf("Error inserting historyOfPurchasing: %s", err.Error())
		return nil, &_historyOfPurchasingException.HistoryOfPurchasingRecordingException{}
	}

	return insertedHistoryOfPurchasing, nil
}

func (r *historyOfPurchasingRepositoryImpl) PlayerHistoryOfPurchasingListing(playerID string) ([]*_historyOfPurchasingEntity.HistoryOfPurchasing, error) {
	historyOfPurchasings := make([]*_historyOfPurchasingEntity.HistoryOfPurchasing, 0)

	if err := r.db.Where("player_id = ?", playerID).Find(&historyOfPurchasings).Error; err != nil {
		r.logger.Errorf("Error finding player historyOfPurchasings: %s", err.Error())
		return nil, &_historyOfPurchasingException.PlayerHistoryOfPurchasingListingException{}
	}

	return historyOfPurchasings, nil
}
