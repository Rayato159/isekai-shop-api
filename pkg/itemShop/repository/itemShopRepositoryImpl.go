package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	entities "github.com/Rayato159/isekai-shop-api/entities"
	_itemShop "github.com/Rayato159/isekai-shop-api/pkg/itemShop/exception"
)

type itemRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemShopRepository {
	return &itemRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *itemRepositoryImpl) TransactionBegin() {
	r.db.Begin()
}

func (r *itemRepositoryImpl) TransactionRollback() {
	r.db.Rollback()
}

func (r *itemRepositoryImpl) TransactionCommit() error {
	return r.db.Commit().Error
}

func (r *itemRepositoryImpl) Listing(itemFilterDto *entities.ItemFilterDto) ([]*entities.Item, error) {
	query := r.db.Model(&entities.Item{})
	if itemFilterDto.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilterDto.Name+"%")
	}
	if itemFilterDto.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilterDto.Description+"%")
	}

	offset := int((itemFilterDto.Page - 1) * itemFilterDto.Size)
	size := int(itemFilterDto.Size)

	items := make([]*entities.Item, 0)

	if err := query.Offset(offset).Limit(size).Find(&items).Error; err != nil {
		r.logger.Error("Failed to find items", err.Error())
		return nil, &_itemShop.ItemListing{}
	}

	return items, nil
}

func (r *itemRepositoryImpl) Counting(itemFilterDto *entities.ItemFilterDto) (int64, error) {
	query := r.db.Model(&entities.Item{}).Where("is_archive = ?", false)

	if itemFilterDto.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilterDto.Name+"%")
	}
	if itemFilterDto.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilterDto.Description+"%")
	}

	var count int64

	if err := query.Count(&count).Error; err != nil {
		r.logger.Error("Counting items failed:", err.Error())
		return -1, &_itemShop.ItemCounting{}
	}

	return count, nil
}

func (r *itemRepositoryImpl) FindByID(itemID uint64) (*entities.Item, error) {
	item := new(entities.Item)

	if err := r.db.First(item, itemID).Error; err != nil {
		r.logger.Error("Finding item failed:", err.Error())
		return nil, &_itemShop.ItemNotFound{ItemID: itemID}
	}

	return item, nil

}

func (r *itemRepositoryImpl) FindByIDList(itemIDs []uint64) ([]*entities.Item, error) {
	items := make([]*entities.Item, 0)

	if err := r.db.Model(&entities.Item{}).Where("id in ?", itemIDs).Find(&items).Error; err != nil {
		r.logger.Error("Finding items by ID failed:", err.Error())
		return nil, &_itemShop.ItemListing{}
	}

	return items, nil
}

func (r *itemRepositoryImpl) PurchaseHistoryRecording(purchasingEntity *entities.PurchaseHistory) (*entities.PurchaseHistory, error) {
	insertedPurchasing := new(entities.PurchaseHistory)

	if err := r.db.Create(purchasingEntity).Scan(insertedPurchasing).Error; err != nil {
		r.logger.Errorf("Purchase history recording failed: %s", err.Error())
		return nil, &_itemShop.HistoryOfPurchaseRecording{}
	}

	return insertedPurchasing, nil
}
