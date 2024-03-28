package service

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"
	_inventoryModel "github.com/Rayato159/isekai-shop-api/pkg/inventory/model"
	_inventoryRepository "github.com/Rayato159/isekai-shop-api/pkg/inventory/repository"
	_itemShopRepository "github.com/Rayato159/isekai-shop-api/pkg/itemShop/repository"
)

type inventoryServiceImpl struct {
	inventoryRepository _inventoryRepository.InventoryRepository
	itemShopRepository  _itemShopRepository.ItemShopRepository
}

func NewInventoryServiceImpl(
	inventoryRepository _inventoryRepository.InventoryRepository,
	itemShopRepository _itemShopRepository.ItemShopRepository,
) InventoryService {
	return &inventoryServiceImpl{
		inventoryRepository,
		itemShopRepository,
	}
}

func (s *inventoryServiceImpl) Listing(playerID string) ([]*_inventoryModel.Inventory, error) {
	inventories, err := s.inventoryRepository.Listing(playerID)
	if err != nil {
		return nil, err
	}

	uniqueItemWithQuantityCounterList := s.getUniqueItemWithQuantityCounterList(inventories)

	return s.buildInventoryListingResult(
		uniqueItemWithQuantityCounterList,
	), nil
}

func (s *inventoryServiceImpl) buildInventoryListingResult(
	uniqueItemWithQuantityCounterList []_inventoryModel.ItemQuantityCounting,
) []*_inventoryModel.Inventory {
	uniqueItemIDList := s.getItemIDList(uniqueItemWithQuantityCounterList)

	itemEntities, err := s.itemShopRepository.FindByIDList(uniqueItemIDList)
	if err != nil {
		return make([]*_inventoryModel.Inventory, 0)
	}

	results := make([]*_inventoryModel.Inventory, 0)
	itemMapWithQuantity := s.getItemMapWithQuantity(uniqueItemWithQuantityCounterList)

	for _, itemEntity := range itemEntities {
		results = append(results, &_inventoryModel.Inventory{
			Item:     itemEntity.ToItemModel(),
			Quantity: itemMapWithQuantity[itemEntity.ID],
		})
	}

	return results
}

func (s *inventoryServiceImpl) getUniqueItemWithQuantityCounterList(
	inventories []*entities.Inventory,
) []_inventoryModel.ItemQuantityCounting {
	itemQuantityCounterList := make([]_inventoryModel.ItemQuantityCounting, 0)

	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range inventories {
		itemMapWithQuantity[inventory.ItemID]++
	}

	for itemID, quantity := range itemMapWithQuantity {
		itemQuantityCounterList = append(itemQuantityCounterList, _inventoryModel.ItemQuantityCounting{
			ItemID:   itemID,
			Quantity: quantity,
		})

	}

	return itemQuantityCounterList
}

func (s *inventoryServiceImpl) getItemIDList(
	uniqueItemWithQuantityCounterList []_inventoryModel.ItemQuantityCounting,
) []uint64 {
	uniqueItemIDList := make([]uint64, 0)

	for _, inventory := range uniqueItemWithQuantityCounterList {
		uniqueItemIDList = append(uniqueItemIDList, inventory.ItemID)
	}

	return uniqueItemIDList
}

func (s *inventoryServiceImpl) getItemMapWithQuantity(
	uniqueItemWithQuantityCounterList []_inventoryModel.ItemQuantityCounting,
) map[uint64]uint {
	itemMapWithQuantity := make(map[uint64]uint)

	for _, itemQuantityCounter := range uniqueItemWithQuantityCounterList {
		itemMapWithQuantity[itemQuantityCounter.ItemID] = itemQuantityCounter.Quantity
	}

	return itemMapWithQuantity
}
