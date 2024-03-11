package repository

import (
	_orderEntity "github.com/Rayato159/isekai-shop-api/modules/order/entity"
	_orderException "github.com/Rayato159/isekai-shop-api/modules/order/exception"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type orderRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewOrderRepositoryImpl(db *gorm.DB, logger echo.Logger) OrderRepository {
	return &orderRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *orderRepositoryImpl) OrderRecording(orderEntity *_orderEntity.Order) (*_orderEntity.Order, error) {
	insertedOrder := new(_orderEntity.Order)

	if err := r.db.Create(orderEntity).Scan(insertedOrder).Error; err != nil {
		r.logger.Errorf("Error inserting order: %s", err.Error())
		return nil, &_orderException.OrderRecordingException{}
	}

	return insertedOrder, nil
}

func (r *orderRepositoryImpl) PlayerOrderListing(playerID string) ([]*_orderEntity.Order, error) {
	orders := make([]*_orderEntity.Order, 0)

	if err := r.db.Where("player_id = ?", playerID).Find(&orders).Error; err != nil {
		r.logger.Errorf("Error finding player orders: %s", err.Error())
		return nil, &_orderException.PlayerOrderListingException{}
	}

	return orders, nil
}
