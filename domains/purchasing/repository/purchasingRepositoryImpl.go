package repository

import (
	_purchasingEntity "github.com/Rayato159/isekai-shop-api/domains/purchasing/entity"
	_purchasingException "github.com/Rayato159/isekai-shop-api/domains/purchasing/exception"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type purchasingRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewPurchasingRepositoryImpl(db *gorm.DB, logger echo.Logger) PurchasingRepository {
	return &purchasingRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *purchasingRepositoryImpl) PurchasingHistoryRecording(purchasingEntity *_purchasingEntity.Purchasing) (*_purchasingEntity.Purchasing, error) {
	insertedPurchasing := new(_purchasingEntity.Purchasing)

	if err := r.db.Create(purchasingEntity).Scan(insertedPurchasing).Error; err != nil {
		r.logger.Errorf("Error inserting purchasing: %s", err.Error())
		return nil, &_purchasingException.PurchasingHistoryRecording{}
	}

	return insertedPurchasing, nil
}
