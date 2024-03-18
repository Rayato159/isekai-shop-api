package service

import (
	"log"

	_balancingException "github.com/Rayato159/isekai-shop-api/domains/balancing/exception"
	_balancingModel "github.com/Rayato159/isekai-shop-api/domains/balancing/model"
	_balancingRepository "github.com/Rayato159/isekai-shop-api/domains/balancing/repository"
	entities "github.com/Rayato159/isekai-shop-api/domains/entities"
	_itemModel "github.com/Rayato159/isekai-shop-api/domains/item/model"
	_itemRepository "github.com/Rayato159/isekai-shop-api/domains/item/repository"
	_playerSource "github.com/Rayato159/isekai-shop-api/domains/player/repository"
	_purchasingRepository "github.com/Rayato159/isekai-shop-api/domains/purchasing/repository"
)

type balancingServiceImpl struct {
	balancingRepository  _balancingRepository.BalancingRepository
	itemRepository       _itemRepository.ItemRepository
	purchasingRepository _purchasingRepository.PurchasingRepository
	inventoryRepository  _playerSource.InventoryRepository
}

func NewBalancingServiceImpl(
	balancingRepository _balancingRepository.BalancingRepository,
	itemRepository _itemRepository.ItemRepository,
	purchasingRepository _purchasingRepository.PurchasingRepository,
	inventoryRepository _playerSource.InventoryRepository,
) BalancingService {
	return &balancingServiceImpl{
		balancingRepository:  balancingRepository,
		itemRepository:       itemRepository,
		purchasingRepository: purchasingRepository,
		inventoryRepository:  inventoryRepository,
	}
}

func (s *balancingServiceImpl) TopUp(topUpReq *_balancingModel.TopUpReq) (*_balancingModel.Balancing, error) {
	balancingEntity := &entities.Balancing{
		PlayerID: topUpReq.PlayerID,
		Amount:   topUpReq.Amount,
	}

	insertedBalancing, err := s.balancingRepository.BalancingRecording(balancingEntity)
	if err != nil {
		return nil, err
	}

	return insertedBalancing.ToBalancingModel(), nil
}

func (s *balancingServiceImpl) PlayerBalanceShowing(playerID string) *_balancingModel.PlayerBalance {
	balanceDto, err := s.balancingRepository.PlayerBalanceShowing(playerID)
	if err != nil {
		return &_balancingModel.PlayerBalance{
			PlayerID: playerID,
			Balance:  0,
		}
	}

	return balanceDto.ToPlayerBalanceModel()
}

// 1. Search item by ID
// 2. Calculate total price
// 3. Check if player has enough balance
// 4. Create purchasing
// 5. Create balancing
// 6. Add item into player inventory
func (s *balancingServiceImpl) ItemBuying(itemBuyingReq *_balancingModel.ItemBuyingReq) (*_balancingModel.Balancing, error) {
	itemEntity, err := s.itemRepository.FindItemByID(itemBuyingReq.ItemID)
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

	insertedBalancing, err := s.balancingRepository.BalancingRecording(&entities.Balancing{
		PlayerID: itemBuyingReq.PlayerID,
		Amount:   -totalPrice,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Balancing entity: %d", insertedBalancing.ID)

	inventory, err := s.inventoryRepository.InventoryFilling(inventoryEntities)
	if err != nil {
		return nil, err
	}
	log.Printf("Inserted inventories: %d", len(inventory))

	return insertedBalancing.ToBalancingModel(), nil
}

// 1. Check if player has enough quantity
// 2. Search item by ID
// 3. Calculate total price
// 4. Create purchasing
// 5. Create balancing
// 6. Delete item into player inventory
func (s *balancingServiceImpl) ItemSelling(itemSellingReq *_balancingModel.ItemSellingReq) (*_balancingModel.Balancing, error) {
	if err := s.checkPlayerItemQuantity(
		itemSellingReq.PlayerID,
		itemSellingReq.ItemID,
		itemSellingReq.Quantity,
	); err != nil {
		return nil, err
	}

	itemEntity, err := s.itemRepository.FindItemByID(itemSellingReq.ItemID)
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

	insertedBalancing, err := s.balancingRepository.BalancingRecording(&entities.Balancing{
		PlayerID: itemSellingReq.PlayerID,
		Amount:   totalPrice,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Balancing entity: %d", insertedBalancing.ID)

	if err := s.inventoryRepository.DeleteItemByLimit(
		itemSellingReq.PlayerID,
		itemSellingReq.ItemID,
		int(itemSellingReq.Quantity),
	); err != nil {
		return nil, err
	}
	log.Printf("Deleted inventories for %d records", itemSellingReq.Quantity)

	return insertedBalancing.ToBalancingModel(), nil
}

func (s *balancingServiceImpl) groupInventoryEntities(itemBuyingReq *_balancingModel.ItemBuyingReq) []*entities.Inventory {
	inventoryEntities := make([]*entities.Inventory, 0)

	for i := 0; i < int(itemBuyingReq.Quantity); i++ {
		inventoryEntities = append(inventoryEntities, &entities.Inventory{
			PlayerID: itemBuyingReq.PlayerID,
			ItemID:   itemBuyingReq.ItemID,
		})
	}

	return inventoryEntities

}

func (s *balancingServiceImpl) checkPlayerItemQuantity(playerID string, itemID uint64, quantity uint) error {
	inventoryCount := s.inventoryRepository.PlayerItemCounting(playerID, itemID)

	if int(inventoryCount) < int(quantity) {
		log.Printf("Player %s has not enough item with id: %d", playerID, itemID)
		return &_balancingException.NotEnoughItemException{ItemID: itemID}
	}

	return nil
}

func (s *balancingServiceImpl) checkPlayerBalance(playerID string, amount int64) error {
	balanceDto, err := s.balancingRepository.PlayerBalanceShowing(playerID)
	if err != nil {
		return err
	}

	if balanceDto.Balance < amount {
		log.Printf("Player %s has not enough balance", playerID)
		return &_balancingException.NotEnoughBalanceException{}
	}

	return nil
}

func (s *balancingServiceImpl) calculateTotalPrice(item *_itemModel.Item, quantity uint) int64 {
	// In a real world scenario, this would be a more complex calculation
	return int64(item.Price) * int64(quantity)
}
