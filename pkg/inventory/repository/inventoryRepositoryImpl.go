package repository

import (
	"github.com/Rayato159/isekai-shop-api/databases"
	entities "github.com/Rayato159/isekai-shop-api/entities"
	_inventory "github.com/Rayato159/isekai-shop-api/pkg/inventory/exception"
	"github.com/labstack/echo/v4"
)

type inventoryRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewInventoryRepositoryImpl(db databases.Database, logger echo.Logger) InventoryRepository {
	return &inventoryRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *inventoryRepositoryImpl) Filling(playerID string, itemID uint64, qty int) ([]*entities.Inventory, error) {
	inventoryEntities := make([]*entities.Inventory, 0)

	for range qty {
		inventoryEntities = append(inventoryEntities, &entities.Inventory{
			PlayerID: playerID,
			ItemID:   itemID,
		})
	}

	if err := r.db.Connect().Create(inventoryEntities).Error; err != nil {
		r.logger.Error("Item creating failed:", err.Error())
		return nil, &_inventory.InventoryFilling{
			PlayerID: playerID,
			ItemID:   itemID,
		}
	}

	return inventoryEntities, nil
}

func (r *inventoryRepositoryImpl) ReverseFilling(inventoryEntities []*entities.Inventory) error {
	for _, inventory := range inventoryEntities {
		inventory.IsDeleted = true

		if err := r.db.Connect().Model(
			&entities.Inventory{},
		).Where(
			"id = ?", inventory.ID,
		).Updates(
			inventory,
		).Error; err != nil {
			r.logger.Error("Removing item failed:", err.Error())
			return &_inventory.PlayerItemRemoving{ItemID: inventory.ID}
		}
	}

	return nil
}

func (r *inventoryRepositoryImpl) Listing(playerID string) ([]*entities.Inventory, error) {
	inventoryEntities := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Where(
		"player_id = ? and is_deleted = ?", playerID, false,
	).Find(&inventoryEntities).Error; err != nil {
		r.logger.Error("Listing player's item failed:", err.Error())
		return nil, &_inventory.PlayerItemsFinding{
			PlayerID: playerID,
		}
	}

	return inventoryEntities, nil
}

func (r *inventoryRepositoryImpl) Removing(playerID string, itemID uint64, limit int) error {
	inventoryEntities, err := r.findPlayerItemInInventoryByID(playerID, itemID, limit)
	if err != nil {
		return err
	}

	for _, inventory := range inventoryEntities {
		inventory.IsDeleted = true

		if err := r.db.Connect().Model(
			&entities.Inventory{},
		).Where(
			"id = ?", inventory.ID,
		).Updates(
			inventory,
		).Error; err != nil {
			r.logger.Error("Removing item failed:", err.Error())
			return &_inventory.PlayerItemRemoving{ItemID: itemID}
		}
	}

	return nil
}

func (r *inventoryRepositoryImpl) ReverseRemoving(playerID string, itemID uint64, limit int) error {
	inventoryEntities, err := r.findPlayerItemInInventoryByID(playerID, itemID, limit)
	if err != nil {
		return err
	}

	for _, inventory := range inventoryEntities {
		inventory.IsDeleted = false

		if err := r.db.Connect().Model(
			&entities.Inventory{},
		).Where(
			"id = ?", inventory.ID,
		).Updates(
			inventory,
		).Error; err != nil {
			r.logger.Error("Removing item failed:", err.Error())
			return &_inventory.PlayerItemRemoving{ItemID: itemID}
		}
	}

	return nil
}

func (r *inventoryRepositoryImpl) PlayerItemCounting(playerID string, itemID uint64) int64 {
	var count int64

	if err := r.db.Connect().Model(
		&entities.Inventory{},
	).Where(
		"player_id = ? and item_id = ? and is_deleted = ?", playerID, itemID, false,
	).Count(&count).Error; err != nil {
		r.logger.Error("Player's item counting failed:", err.Error())
		return -1
	}

	return count
}

func (r *inventoryRepositoryImpl) findPlayerItemInInventoryByID(
	playerID string,
	itemID uint64,
	limit int,
) ([]*entities.Inventory, error) {
	inventoryEntities := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Where(
		"player_id = ? and item_id = ? and is_deleted = ?", playerID, itemID, false,
	).Limit(
		limit,
	).Find(&inventoryEntities).Error; err != nil {
		r.logger.Error("Finding player's item in inventory failed:", err.Error())
		return nil, &_inventory.PlayerItemsFinding{PlayerID: playerID}
	}

	return inventoryEntities, nil
}
