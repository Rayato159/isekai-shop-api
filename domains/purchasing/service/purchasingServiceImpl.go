package service

import (
	"log"

	_inventoryRepository "github.com/Rayato159/isekai-shop-api/domains/inventory/repository"
	_itemGettingModel "github.com/Rayato159/isekai-shop-api/domains/itemGetting/model"
	_itemGettingRepository "github.com/Rayato159/isekai-shop-api/domains/itemGetting/repository"
	_playerBalancingModel "github.com/Rayato159/isekai-shop-api/domains/playerBalancing/model"
	_playerBalancingRepository "github.com/Rayato159/isekai-shop-api/domains/playerBalancing/repository"
	_purchasingException "github.com/Rayato159/isekai-shop-api/domains/purchasing/exception"
	_purchasingModel "github.com/Rayato159/isekai-shop-api/domains/purchasing/model"
	_purchasingRepository "github.com/Rayato159/isekai-shop-api/domains/purchasing/repository"

	entities "github.com/Rayato159/isekai-shop-api/entities"
)

type purchasingServiceImpl struct {
	playerBalancingRepository _playerBalancingRepository.PlayerBalancingRepository
	itemRepository            _itemGettingRepository.ItemGettingRepository
	purchasingRepository      _purchasingRepository.PurchasingRepository
	inventoryRepository       _inventoryRepository.InventoryRepository
}

func NewPurchasingServiceImpl(
	playerBalancingRepository _playerBalancingRepository.PlayerBalancingRepository,
	itemRepository _itemGettingRepository.ItemGettingRepository,
	purchasingRepository _purchasingRepository.PurchasingRepository,
	inventoryRepository _inventoryRepository.InventoryRepository,
) PurchasingService {
	return &purchasingServiceImpl{
		playerBalancingRepository,
		itemRepository,
		purchasingRepository,
		inventoryRepository,
	}
}

// 1. Search item by ID
// 2. Calculate total price
// 3. Check if player has enough balance
// 4. Create purchasing
// 5. Create playerBalancing
// 6. Add item into player inventory
func (s *purchasingServiceImpl) ItemBuying(itemBuyingReq *_purchasingModel.ItemBuyingReq) (*_playerBalancingModel.PlayerBalancing, error) {
	itemEntity, err := s.itemRepository.FindByID(itemBuyingReq.ItemID)
	if err != nil {
		return nil, err
	}

	totalPrice := s.calculateTotalPrice(itemEntity.ToItemModel(), itemBuyingReq.Quantity)

	if err := s.checkPlayerBalance(itemBuyingReq.PlayerID, totalPrice); err != nil {
		return nil, err
	}

	insertedPurchasing, err := s.purchasingRepository.PurchasingHistoryRecording(&entities.PurchasingHistory{
		PlayerID:        itemBuyingReq.PlayerID,
		ItemID:          itemEntity.ID,
		ItemName:        itemEntity.Name,
		ItemDescription: itemEntity.Description,
		ItemPrice:       itemEntity.Price,
		ItemPicture:     itemEntity.Picture,
		Quantity:        itemBuyingReq.Quantity,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Inserted purchasing: %d", insertedPurchasing.ID)

	inventoryEntities := s.groupInventoryEntities(itemBuyingReq)

	insertedBalancing, err := s.playerBalancingRepository.Recording(&entities.PlayerBalancing{
		PlayerID: itemBuyingReq.PlayerID,
		Amount:   -totalPrice,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Balancing entity: %d", insertedBalancing.ID)

	inventory, err := s.inventoryRepository.Filling(inventoryEntities)
	if err != nil {
		return nil, err
	}
	log.Printf("Inserted inventories: %d", len(inventory))

	return insertedBalancing.ToPlayerBalancingModel(), nil
}

// 1. Check if player has enough quantity
// 2. Search item by ID
// 3. Calculate total price
// 4. Create purchasing
// 5. Create playerBalancing
// 6. Delete item into player inventory
func (s *purchasingServiceImpl) ItemSelling(itemSellingReq *_purchasingModel.ItemSellingReq) (*_playerBalancingModel.PlayerBalancing, error) {
	if err := s.checkPlayerItemQuantity(
		itemSellingReq.PlayerID,
		itemSellingReq.ItemID,
		itemSellingReq.Quantity,
	); err != nil {
		return nil, err
	}

	itemEntity, err := s.itemRepository.FindByID(itemSellingReq.ItemID)
	if err != nil {
		return nil, err
	}

	totalPrice := s.calculateTotalPrice(itemEntity.ToItemModel(), itemSellingReq.Quantity)
	totalPrice = totalPrice / 2

	insertedPurchasing, err := s.purchasingRepository.PurchasingHistoryRecording(&entities.PurchasingHistory{
		PlayerID:        itemSellingReq.PlayerID,
		ItemID:          itemEntity.ID,
		ItemName:        itemEntity.Name,
		ItemDescription: itemEntity.Description,
		ItemPrice:       itemEntity.Price,
		ItemPicture:     itemEntity.Picture,
		Quantity:        itemSellingReq.Quantity,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Inserted purchasing: %d", insertedPurchasing.ID)

	insertedBalancing, err := s.playerBalancingRepository.Recording(&entities.PlayerBalancing{
		PlayerID: itemSellingReq.PlayerID,
		Amount:   totalPrice,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Balancing entity: %d", insertedBalancing.ID)

	if err := s.inventoryRepository.DeletePlayerItemByLimit(
		itemSellingReq.PlayerID,
		itemSellingReq.ItemID,
		int(itemSellingReq.Quantity),
	); err != nil {
		return nil, err
	}
	log.Printf("Deleted inventories for %d records", itemSellingReq.Quantity)

	return insertedBalancing.ToPlayerBalancingModel(), nil
}

func (s *purchasingServiceImpl) groupInventoryEntities(itemBuyingReq *_purchasingModel.ItemBuyingReq) []*entities.Inventory {
	inventoryEntities := make([]*entities.Inventory, 0)

	for i := 0; i < int(itemBuyingReq.Quantity); i++ {
		inventoryEntities = append(inventoryEntities, &entities.Inventory{
			PlayerID: itemBuyingReq.PlayerID,
			ItemID:   itemBuyingReq.ItemID,
		})
	}

	return inventoryEntities

}

func (s *purchasingServiceImpl) checkPlayerItemQuantity(playerID string, itemID uint64, quantity uint) error {
	inventoryCount := s.inventoryRepository.PlayerItemCounting(playerID, itemID)

	if int(inventoryCount) < int(quantity) {
		log.Printf("Player %s has not enough item with id: %d", playerID, itemID)
		return &_purchasingException.NotEnoughItemException{ItemID: itemID}
	}

	return nil
}

func (s *purchasingServiceImpl) checkPlayerBalance(playerID string, amount int64) error {
	balanceDto, err := s.playerBalancingRepository.Showing(playerID)
	if err != nil {
		return err
	}

	if balanceDto.Balance < amount {
		log.Printf("Player %s has not enough balance", playerID)
		return &_purchasingException.NotEnoughBalanceException{}
	}

	return nil
}

func (s *purchasingServiceImpl) calculateTotalPrice(item *_itemGettingModel.Item, quantity uint) int64 {
	// In a real world scenario, this would be a more complex calculation
	return int64(item.Price) * int64(quantity)
}