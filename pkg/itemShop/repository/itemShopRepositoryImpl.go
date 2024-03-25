package repository

import (
	"github.com/labstack/echo/v4"

	"github.com/Rayato159/isekai-shop-api/databases"
	entities "github.com/Rayato159/isekai-shop-api/entities"
	_itemShop "github.com/Rayato159/isekai-shop-api/pkg/itemShop/exception"
	_itemShopModel "github.com/Rayato159/isekai-shop-api/pkg/itemShop/model"
)

type itemRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db databases.Database, logger echo.Logger) ItemShopRepository {
	return &itemRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *itemRepositoryImpl) TransactionBegin() {
	r.db.Connect().Begin()
}

func (r *itemRepositoryImpl) TransactionRollback() {
	r.db.Connect().Rollback()
}

func (r *itemRepositoryImpl) TransactionCommit() error {
	return r.db.Connect().Commit().Error
}

func (r *itemRepositoryImpl) Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) {
	query := r.db.Connect().Model(&entities.Item{}).Where("is_archive = ?", false)

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}
	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	offset := int((itemFilter.Page - 1) * itemFilter.Size)
	size := int(itemFilter.Size)

	items := make([]*entities.Item, 0)

	if err := query.Offset(offset).Limit(size).Find(&items).Order("id asc").Error; err != nil {
		r.logger.Error("Failed to find items", err.Error())
		return nil, &_itemShop.ItemListing{}
	}

	return items, nil
}

func (r *itemRepositoryImpl) Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error) {
	query := r.db.Connect().Model(&entities.Item{}).Where("is_archive = ?", false)

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}
	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
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

	if err := r.db.Connect().First(item, itemID).Error; err != nil {
		r.logger.Error("Finding item failed:", err.Error())
		return nil, &_itemShop.ItemNotFound{ItemID: itemID}
	}

	return item, nil

}

func (r *itemRepositoryImpl) FindByIDList(itemIDs []uint64) ([]*entities.Item, error) {
	items := make([]*entities.Item, 0)

	if err := r.db.Connect().Model(&entities.Item{}).Where("id in ?", itemIDs).Find(&items).Error; err != nil {
		r.logger.Error("Finding items by ID failed:", err.Error())
		return nil, &_itemShop.ItemListing{}
	}

	return items, nil
}

func (r *itemRepositoryImpl) PurchaseHistoryRecording(purchasingEntity *entities.PurchaseHistory) (*entities.PurchaseHistory, error) {
	insertedPurchasing := new(entities.PurchaseHistory)

	if err := r.db.Connect().Create(purchasingEntity).Scan(insertedPurchasing).Error; err != nil {
		r.logger.Errorf("Purchase history recording failed: %s", err.Error())
		return nil, &_itemShop.HistoryOfPurchaseRecording{}
	}

	return insertedPurchasing, nil
}
