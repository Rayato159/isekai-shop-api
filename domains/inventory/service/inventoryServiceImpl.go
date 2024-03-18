package service

import (
	_inventoryModel "github.com/Rayato159/isekai-shop-api/domains/inventory/model"
	_inventory "github.com/Rayato159/isekai-shop-api/domains/inventory/repository"
	_inventoryRepository "github.com/Rayato159/isekai-shop-api/domains/inventory/repository"
	_itemModel "github.com/Rayato159/isekai-shop-api/domains/itemGetting/model"
	_itemGettingRepository "github.com/Rayato159/isekai-shop-api/domains/itemGetting/repository"
	_playerRepository "github.com/Rayato159/isekai-shop-api/domains/player/repository"
	entities "github.com/Rayato159/isekai-shop-api/entities"
)

type inventoryImpl struct {
	playerRepository      _playerRepository.PlayerRepository
	inventoryRepository   _inventoryRepository.InventoryRepository
	itemGettingRepository _itemGettingRepository.ItemGettingRepository
}

func NewInventoryServiceImpl(
	playerRepository _playerRepository.PlayerRepository,
	inventoryRepository _inventory.InventoryRepository,
	itemGettingRepository _itemGettingRepository.ItemGettingRepository,
) InventoryService {
	return &inventoryImpl{
		playerRepository,
		inventoryRepository,
		itemGettingRepository,
	}
}

func (s *inventoryImpl) Listing(playerID string) ([]*_inventoryModel.Inventory, error) {
	inventories, err := s.inventoryRepository.Listing(playerID)
	if err != nil {
		return nil, err
	}

	uniqueItemWithQuantityCounterList := s.getUniqueItemWithQuantityCounterList(inventories)

	return s.buildInventoryListingResult(
		uniqueItemWithQuantityCounterList,
	), nil
}

func (s *inventoryImpl) buildInventoryListingResult(
	uniqueItemWithQuantityCounterList []_inventoryModel.ItemQuantityCounting,
) []*_inventoryModel.Inventory {
	uniqueItemIDList := s.getItemIDList(uniqueItemWithQuantityCounterList)

	itemEntities, err := s.itemGettingRepository.FindByIDList(uniqueItemIDList)
	if err != nil {
		return make([]*_inventoryModel.Inventory, 0)
	}

	results := make([]*_inventoryModel.Inventory, 0)
	itemMapWithQuantity := s.getItemMapWithQuantity(uniqueItemWithQuantityCounterList)

	for _, itemEntity := range itemEntities {
		results = append(results, &_inventoryModel.Inventory{
			Item: &_itemModel.Item{
				ID:          itemEntity.ID,
				Name:        itemEntity.Name,
				Description: itemEntity.Description,
				Picture:     itemEntity.Picture,
				Price:       itemEntity.Price,
			},
			Quantity: itemMapWithQuantity[itemEntity.ID],
		})

	}

	return results
}

func (s *inventoryImpl) getUniqueItemWithQuantityCounterList(
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

func (s *inventoryImpl) getItemIDList(
	uniqueItemWithQuantityCounterList []_inventoryModel.ItemQuantityCounting,
) []uint64 {
	uniqueItemIDList := make([]uint64, 0)

	for _, inventory := range uniqueItemWithQuantityCounterList {
		uniqueItemIDList = append(uniqueItemIDList, inventory.ItemID)
	}

	return uniqueItemIDList
}

func (s *inventoryImpl) getItemMapWithQuantity(
	uniqueItemWithQuantityCounterList []_inventoryModel.ItemQuantityCounting,
) map[uint64]uint {
	itemMapWithQuantity := make(map[uint64]uint)

	for _, itemQuantityCounter := range uniqueItemWithQuantityCounterList {
		itemMapWithQuantity[itemQuantityCounter.ItemID] = itemQuantityCounter.Quantity
	}

	return itemMapWithQuantity
}
