package service

import (
	entities "github.com/Rayato159/isekai-shop-api/domains/entities"
	_itemModel "github.com/Rayato159/isekai-shop-api/domains/item/model"
	_itemRepository "github.com/Rayato159/isekai-shop-api/domains/item/repository"
	_playerModel "github.com/Rayato159/isekai-shop-api/domains/player/model"
	_playerSource "github.com/Rayato159/isekai-shop-api/domains/player/repository"
)

type playerServiceImpl struct {
	playerRepository    _playerSource.PlayerRepository
	inventoryRepository _playerSource.InventoryRepository
	itemRepository      _itemRepository.ItemRepository
}

func NewPlayerServiceImpl(
	playerRepository _playerSource.PlayerRepository,
	inventoryRepository _playerSource.InventoryRepository,
	itemRepository _itemRepository.ItemRepository,
) PlayerService {
	return &playerServiceImpl{
		playerRepository,
		inventoryRepository,
		itemRepository,
	}
}

func (s *playerServiceImpl) PlayerInventoryListing(playerID string) ([]*_playerModel.Inventory, error) {
	inventories, err := s.inventoryRepository.InventorySearching(playerID)
	if err != nil {
		return nil, err
	}

	uniqueItemWithQuantityCounterList := s.getUniqueItemWithQuantityCounterList(inventories)

	return s.buildInventoryListingResult(
		uniqueItemWithQuantityCounterList,
	), nil
}

func (s *playerServiceImpl) buildInventoryListingResult(
	uniqueItemWithQuantityCounterList []_playerModel.ItemQuantityCounting,
) []*_playerModel.Inventory {
	uniqueItemIDList := s.getItemIDList(uniqueItemWithQuantityCounterList)

	itemEntities, err := s.itemRepository.FindItemByIDs(uniqueItemIDList)
	if err != nil {
		return make([]*_playerModel.Inventory, 0)
	}

	results := make([]*_playerModel.Inventory, 0)
	itemMapWithQuantity := s.getItemMapWithQuantity(uniqueItemWithQuantityCounterList)

	for _, itemEntity := range itemEntities {
		results = append(results, &_playerModel.Inventory{
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

func (s *playerServiceImpl) getUniqueItemWithQuantityCounterList(
	inventories []*entities.Inventory,
) []_playerModel.ItemQuantityCounting {
	itemQuantityCounterList := make([]_playerModel.ItemQuantityCounting, 0)

	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range inventories {
		itemMapWithQuantity[inventory.ItemID]++
	}

	for itemID, quantity := range itemMapWithQuantity {
		itemQuantityCounterList = append(itemQuantityCounterList, _playerModel.ItemQuantityCounting{
			ItemID:   itemID,
			Quantity: quantity,
		})

	}

	return itemQuantityCounterList
}

func (s *playerServiceImpl) getItemIDList(
	uniqueItemWithQuantityCounterList []_playerModel.ItemQuantityCounting,
) []uint64 {
	uniqueItemIDList := make([]uint64, 0)

	for _, inventory := range uniqueItemWithQuantityCounterList {
		uniqueItemIDList = append(uniqueItemIDList, inventory.ItemID)
	}

	return uniqueItemIDList
}

func (s *playerServiceImpl) getItemMapWithQuantity(
	uniqueItemWithQuantityCounterList []_playerModel.ItemQuantityCounting,
) map[uint64]uint {
	itemMapWithQuantity := make(map[uint64]uint)

	for _, itemQuantityCounter := range uniqueItemWithQuantityCounterList {
		itemMapWithQuantity[itemQuantityCounter.ItemID] = itemQuantityCounter.Quantity
	}

	return itemMapWithQuantity
}
