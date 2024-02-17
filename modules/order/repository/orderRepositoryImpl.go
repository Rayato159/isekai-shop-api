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

func NewOrderRepository(db *gorm.DB, logger echo.Logger) OrderRepository {
	return &orderRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *orderRepositoryImpl) InsertOrder(orderEntity *_orderEntity.Order) (*_orderEntity.Order, error) {
	insertedOrder := new(_orderEntity.Order)

	if err := r.db.Create(orderEntity).Scan(insertedOrder).Error; err != nil {
		r.logger.Errorf("Error inserting order: %s", err.Error())
		return nil, &_orderException.InsertOrderException{}
	}

	return insertedOrder, nil
}
