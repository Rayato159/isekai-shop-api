package repository

import (
	"github.com/Rayato159/isekai-shop-api/databases"
	entities "github.com/Rayato159/isekai-shop-api/entities"
	_itemShopModel "github.com/Rayato159/isekai-shop-api/pkg/itemShop/model"
)

type ItemShopRepository interface {
	Listing(itemFilterDto *_itemShopModel.ItemFilter) ([]*entities.Item, error)
	FindByID(itemID uint64) (*entities.Item, error)
	FindByIDList(itemIDs []uint64) ([]*entities.Item, error)
	Counting(itemFilterDto *_itemShopModel.ItemFilter) (int64, error)
	PurchaseHistoryRecording(purchasingEntity *entities.PurchaseHistory) (*entities.PurchaseHistory, error)
	Transaction() databases.Database
}
