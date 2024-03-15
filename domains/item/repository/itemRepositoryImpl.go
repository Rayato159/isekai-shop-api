package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_itemEntity "github.com/Rayato159/isekai-shop-api/domains/item/entity"
	_itemException "github.com/Rayato159/isekai-shop-api/domains/item/exception"
)

type itemRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemRepository {
	return &itemRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *itemRepositoryImpl) ItemListing(itemFilterDto *_itemEntity.ItemFilterDto) ([]*_itemEntity.Item, error) {
	query := r.db.Model(&_itemEntity.Item{})
	if itemFilterDto.Name != "" {
		query = query.Where("name LIKE ?", "%"+itemFilterDto.Name+"%")
	}
	if itemFilterDto.Description != "" {
		query = query.Where("description LIKE ?", "%"+itemFilterDto.Description+"%")
	}

	offset := int((itemFilterDto.Page - 1) * itemFilterDto.Size)
	size := int(itemFilterDto.Size)

	items := make([]*_itemEntity.Item, 0)

	if err := query.Offset(offset).Limit(size).Find(&items).Error; err != nil {
		r.logger.Error("Failed to find items", err.Error())
		return nil, &_itemException.ItemListingException{}
	}

	return items, nil
}

func (r *itemRepositoryImpl) ItemCounting(itemFilterDto *_itemEntity.ItemFilterDto) (int64, error) {
	query := r.db.Model(&_itemEntity.Item{}).Where("is_archive = ?", false)

	if itemFilterDto.Name != "" {
		query = query.Where("name LIKE ?", "%"+itemFilterDto.Name+"%")
	}
	if itemFilterDto.Description != "" {
		query = query.Where("description LIKE ?", "%"+itemFilterDto.Description+"%")
	}

	var count int64

	if err := query.Count(&count).Error; err != nil {
		r.logger.Error("Failed to count items", err.Error())
		return -1, &_itemException.ItemCountingException{}
	}

	return count, nil
}

func (r *itemRepositoryImpl) FindItemByID(itemID uint64) (*_itemEntity.Item, error) {
	item := new(_itemEntity.Item)

	if err := r.db.First(item, itemID).Error; err != nil {
		r.logger.Error("Failed to find item", err.Error())
		return nil, &_itemException.ItemNotFoundException{ItemID: itemID}
	}

	return item, nil

}

func (r *itemRepositoryImpl) ItemCreating(itemEntity *_itemEntity.Item) (*_itemEntity.Item, error) {
	insertedItem := new(_itemEntity.Item)

	if err := r.db.Create(itemEntity).Scan(insertedItem).Error; err != nil {
		r.logger.Error("Failed to insert item", err.Error())
		return nil, &_itemException.ItemCreatingException{}
	}

	return insertedItem, nil
}

func (r *itemRepositoryImpl) ItemEditing(itemID uint64, updateItemDto *_itemEntity.ItemEditingDto) (uint64, error) {
	tx := r.db.Model(&_itemEntity.Item{}).Where(
		"id = ?", itemID,
	).Updates(
		updateItemDto,
	)

	if tx.Error != nil {
		r.logger.Error("Failed to update item", tx.Error.Error())
		return 0, &_itemException.ItemEditingException{}
	}

	return itemID, nil
}

func (r *itemRepositoryImpl) ItemArchiving(itemID uint64) error {
	if err := r.db.Table("items").Where(
		"id = ?", itemID,
	).Update(
		"is_archive", true,
	).Error; err != nil {
		r.logger.Error("Failed to archive item", err.Error())
		return &_itemException.ItemArchivingException{ItemID: itemID}
	}

	return nil
}

func (r *itemRepositoryImpl) FindItemByIDs(itemIDs []uint64) ([]*_itemEntity.Item, error) {
	items := make([]*_itemEntity.Item, 0)

	if err := r.db.Model(&_itemEntity.Item{}).Where("id IN ?", itemIDs).Find(&items).Error; err != nil {
		r.logger.Error("Failed to find items by IDs", err.Error())
		return nil, &_itemException.ItemListingException{}
	}

	return items, nil
}
