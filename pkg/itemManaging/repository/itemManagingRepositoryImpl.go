package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"
	_itemManaging "github.com/Rayato159/isekai-shop-api/pkg/itemManaging/exception"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type itemMangingRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemManagingRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemManagingRepository {
	return &itemMangingRepositoryImpl{db, logger}
}

func (r *itemMangingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	insertedItem := new(entities.Item)

	if err := r.db.Create(itemEntity).Scan(insertedItem).Error; err != nil {
		r.logger.Error("Item creating failed:", err.Error())
		return nil, &_itemManaging.ItemCreating{}
	}

	return insertedItem, nil
}

func (r *itemMangingRepositoryImpl) Editing(itemID uint64, updateItemDto *entities.ItemEditingDto) (uint64, error) {
	tx := r.db.Model(&entities.Item{}).Where(
		"id = ?", itemID,
	).Updates(
		updateItemDto,
	)

	if tx.Error != nil {
		r.logger.Error("Editing item failed:", tx.Error.Error())
		return 0, &_itemManaging.ItemEditing{}
	}

	return itemID, nil
}

func (r *itemMangingRepositoryImpl) Archiving(itemID uint64) error {
	if err := r.db.Table("items").Where(
		"id = ?", itemID,
	).Update(
		"is_archive", true,
	).Error; err != nil {
		r.logger.Error("Archiving item failed:", err.Error())
		return &_itemManaging.ItemArchiving{ItemID: itemID}
	}

	return nil
}
