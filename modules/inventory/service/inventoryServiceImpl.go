package service

import (
	_inventoryEntity "github.com/Rayato159/isekai-shop-api/modules/inventory/entity"
	_inventoryModel "github.com/Rayato159/isekai-shop-api/modules/inventory/model"
	_inventoryRepository "github.com/Rayato159/isekai-shop-api/modules/inventory/repository"
	_itemModel "github.com/Rayato159/isekai-shop-api/modules/item/model"
	_itemRepository "github.com/Rayato159/isekai-shop-api/modules/item/repository"
	"github.com/labstack/echo/v4"
)

type inventoryServiceImpl struct {
	inventoryRepository _inventoryRepository.InventoryRepository
	itemRepository      _itemRepository.ItemRepository
	logger              echo.Logger
}

func NewInventoryService(
	inventoryRepository _inventoryRepository.InventoryRepository,
	itemRepository _itemRepository.ItemRepository,
	logger echo.Logger,
) InventoryService {
	return &inventoryServiceImpl{
		inventoryRepository: inventoryRepository,
		itemRepository:      itemRepository,
		logger:              logger,
	}
}

func (s *inventoryServiceImpl) PlayerInventoryListing(playerID string) ([]*_inventoryModel.Inventory, error) {
	inventories, err := s.inventoryRepository.FindPlayerInventories(playerID)
	if err != nil {
		return nil, err
	}

	uniqueItemWithQuantityCounterList := s.getUniqueItemWithQuantityCounterList(inventories)

	return s.buildInventoryListingResult(
		inventories,
		uniqueItemWithQuantityCounterList,
	), nil
}

func (s *inventoryServiceImpl) buildInventoryListingResult(
	inventories []*_inventoryEntity.Inventory,
	uniqueItemWithQuantityCounterList []_inventoryModel.ItemQuantityCounter,
) []*_inventoryModel.Inventory {
	uniqueItemIDList := s.getItemIDList(uniqueItemWithQuantityCounterList)

	itemEntities, err := s.itemRepository.FindItemByIDs(uniqueItemIDList)
	if err != nil {
		s.logger.Error("Failed to find items", err.Error())
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

func (s *inventoryServiceImpl) getUniqueItemWithQuantityCounterList(
	inventories []*_inventoryEntity.Inventory,
) []_inventoryModel.ItemQuantityCounter {
	itemQuantityCounterList := make([]_inventoryModel.ItemQuantityCounter, 0)

	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range inventories {
		itemMapWithQuantity[inventory.ItemID]++
	}

	for itemID, quantity := range itemMapWithQuantity {
		itemQuantityCounterList = append(itemQuantityCounterList, _inventoryModel.ItemQuantityCounter{
			ItemID:   itemID,
			Quantity: quantity,
		})

	}

	return itemQuantityCounterList
}

func (s *inventoryServiceImpl) getItemIDList(
	uniqueItemWithQuantityCounterList []_inventoryModel.ItemQuantityCounter,
) []uint64 {
	uniqueItemIDList := make([]uint64, 0)

	for _, inventory := range uniqueItemWithQuantityCounterList {
		uniqueItemIDList = append(uniqueItemIDList, inventory.ItemID)
	}

	return uniqueItemIDList
}

func (s *inventoryServiceImpl) getItemMapWithQuantity(
	uniqueItemWithQuantityCounterList []_inventoryModel.ItemQuantityCounter,
) map[uint64]uint {
	itemMapWithQuantity := make(map[uint64]uint)

	for _, itemQuantityCounter := range uniqueItemWithQuantityCounterList {
		itemMapWithQuantity[itemQuantityCounter.ItemID] = itemQuantityCounter.Quantity
	}

	return itemMapWithQuantity
}
