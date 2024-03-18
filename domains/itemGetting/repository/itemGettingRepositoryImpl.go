package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_itemGettingException "github.com/Rayato159/isekai-shop-api/domains/itemGetting/exception"
	entities "github.com/Rayato159/isekai-shop-api/entities"
)

type itemRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemGettingRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemGettingRepository {
	return &itemRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *itemRepositoryImpl) Listing(itemFilterDto *entities.ItemFilterDto) ([]*entities.Item, error) {
	query := r.db.Model(&entities.Item{})
	if itemFilterDto.Name != "" {
		query = query.Where("name LIKE ?", "%"+itemFilterDto.Name+"%")
	}
	if itemFilterDto.Description != "" {
		query = query.Where("description LIKE ?", "%"+itemFilterDto.Description+"%")
	}

	offset := int((itemFilterDto.Page - 1) * itemFilterDto.Size)
	size := int(itemFilterDto.Size)

	items := make([]*entities.Item, 0)

	if err := query.Offset(offset).Limit(size).Find(&items).Error; err != nil {
		r.logger.Error("Failed to find items", err.Error())
		return nil, &_itemGettingException.ItemListingException{}
	}

	return items, nil
}

func (r *itemRepositoryImpl) Counting(itemFilterDto *entities.ItemFilterDto) (int64, error) {
	query := r.db.Model(&entities.Item{}).Where("is_archive = ?", false)

	if itemFilterDto.Name != "" {
		query = query.Where("name LIKE ?", "%"+itemFilterDto.Name+"%")
	}
	if itemFilterDto.Description != "" {
		query = query.Where("description LIKE ?", "%"+itemFilterDto.Description+"%")
	}

	var count int64

	if err := query.Count(&count).Error; err != nil {
		r.logger.Error("Failed to count items", err.Error())
		return -1, &_itemGettingException.ItemCountingException{}
	}

	return count, nil
}

func (r *itemRepositoryImpl) FindByID(itemID uint64) (*entities.Item, error) {
	item := new(entities.Item)

	if err := r.db.First(item, itemID).Error; err != nil {
		r.logger.Error("Failed to find item", err.Error())
		return nil, &_itemGettingException.ItemNotFoundException{ItemID: itemID}
	}

	return item, nil

}

func (r *itemRepositoryImpl) FindByIDList(itemIDs []uint64) ([]*entities.Item, error) {
	items := make([]*entities.Item, 0)

	if err := r.db.Model(&entities.Item{}).Where("id IN ?", itemIDs).Find(&items).Error; err != nil {
		r.logger.Error("Failed to find items by IDs", err.Error())
		return nil, &_itemGettingException.ItemListingException{}
	}

	return items, nil
}
