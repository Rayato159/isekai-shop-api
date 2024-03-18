package service

import (
	"log"

	_inventoryRepository "github.com/Rayato159/isekai-shop-api/domains/inventory/repository"
	_itemShopException "github.com/Rayato159/isekai-shop-api/domains/itemShop/exception"
	_itemShopModel "github.com/Rayato159/isekai-shop-api/domains/itemShop/model"
	_itemShopRepository "github.com/Rayato159/isekai-shop-api/domains/itemShop/repository"
	_playerBalancingModel "github.com/Rayato159/isekai-shop-api/domains/playerBalancing/model"
	_playerBalancingRepository "github.com/Rayato159/isekai-shop-api/domains/playerBalancing/repository"
	entities "github.com/Rayato159/isekai-shop-api/entities"
)

type itemShopServiceImpl struct {
	itemShopRepository        _itemShopRepository.ItemShopRepository
	playerBalancingRepository _playerBalancingRepository.PlayerBalancingRepository
	inventoryRepository       _inventoryRepository.InventoryRepository
}

func NewItemShopServiceImpl(
	itemShopRepository _itemShopRepository.ItemShopRepository,
	playerBalancingRepository _playerBalancingRepository.PlayerBalancingRepository,
	inventoryRepository _inventoryRepository.InventoryRepository,
) ItemShopService {
	return &itemShopServiceImpl{
		itemShopRepository,
		playerBalancingRepository,
		inventoryRepository,
	}
}

func (s *itemShopServiceImpl) Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error) {
	itemFilterDto := &entities.ItemFilterDto{
		Name:        itemFilter.Name,
		Description: itemFilter.Description,
		PaginateDto: entities.PaginateDto{
			Page: itemFilter.Page,
			Size: itemFilter.Size,
		},
	}

	itemEntityList, err := s.itemShopRepository.Listing(itemFilterDto)
	if err != nil {
		return nil, err
	}

	totalItems, err := s.itemShopRepository.Counting(itemFilterDto)
	if err != nil {
		return nil, err
	}

	size := itemFilter.Paginate.Size
	page := itemFilter.Paginate.Page
	totalPage := s.totalPageCalculation(totalItems, size)

	result := s.buildItemResultsResponse(itemEntityList, page, totalPage)

	return result, nil
}

// 1. Search item by ID
// 2. Calculate total price
// 3. Check if player has enough balance
// 4. Create itemShop
// 5. Create playerBalancing
// 6. Add item into player inventory
func (s *itemShopServiceImpl) Buying(buyingReq *_itemShopModel.BuyingReq) (*_playerBalancingModel.PlayerBalancing, error) {
	itemEntity, err := s.itemShopRepository.FindByID(buyingReq.ItemID)
	if err != nil {
		return nil, err
	}

	totalPrice := s.calculateTotalPrice(itemEntity.ToItemModel(), buyingReq.Quantity)

	if err := s.checkPlayerBalance(buyingReq.PlayerID, totalPrice); err != nil {
		return nil, err
	}

	insertedPurchasing, err := s.itemShopRepository.PurchasingHistoryRecording(&entities.PurchasingHistory{
		PlayerID:        buyingReq.PlayerID,
		ItemID:          itemEntity.ID,
		ItemName:        itemEntity.Name,
		ItemDescription: itemEntity.Description,
		ItemPrice:       itemEntity.Price,
		ItemPicture:     itemEntity.Picture,
		Quantity:        buyingReq.Quantity,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Inserted itemShop: %d", insertedPurchasing.ID)

	inventoryEntities := s.groupInventoryEntities(buyingReq)

	insertedBalancing, err := s.playerBalancingRepository.Recording(&entities.PlayerBalancing{
		PlayerID: buyingReq.PlayerID,
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
// 4. Create itemShop
// 5. Create playerBalancing
// 6. Delete item into player inventory
func (s *itemShopServiceImpl) Selling(sellingReq *_itemShopModel.SellingReq) (*_playerBalancingModel.PlayerBalancing, error) {
	if err := s.checkPlayerItemQuantity(
		sellingReq.PlayerID,
		sellingReq.ItemID,
		sellingReq.Quantity,
	); err != nil {
		return nil, err
	}

	itemEntity, err := s.itemShopRepository.FindByID(sellingReq.ItemID)
	if err != nil {
		return nil, err
	}

	totalPrice := s.calculateTotalPrice(itemEntity.ToItemModel(), sellingReq.Quantity)
	totalPrice = totalPrice / 2

	insertedPurchasing, err := s.itemShopRepository.PurchasingHistoryRecording(&entities.PurchasingHistory{
		PlayerID:        sellingReq.PlayerID,
		ItemID:          itemEntity.ID,
		ItemName:        itemEntity.Name,
		ItemDescription: itemEntity.Description,
		ItemPrice:       itemEntity.Price,
		ItemPicture:     itemEntity.Picture,
		Quantity:        sellingReq.Quantity,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Inserted itemShop: %d", insertedPurchasing.ID)

	insertedBalancing, err := s.playerBalancingRepository.Recording(&entities.PlayerBalancing{
		PlayerID: sellingReq.PlayerID,
		Amount:   totalPrice,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Balancing entity: %d", insertedBalancing.ID)

	if err := s.inventoryRepository.DeletePlayerItemByLimit(
		sellingReq.PlayerID,
		sellingReq.ItemID,
		int(sellingReq.Quantity),
	); err != nil {
		return nil, err
	}
	log.Printf("Deleted inventories for %d records", sellingReq.Quantity)

	return insertedBalancing.ToPlayerBalancingModel(), nil
}

func (s *itemShopServiceImpl) groupInventoryEntities(buyingReq *_itemShopModel.BuyingReq) []*entities.Inventory {
	inventoryEntities := make([]*entities.Inventory, 0)

	for i := 0; i < int(buyingReq.Quantity); i++ {
		inventoryEntities = append(inventoryEntities, &entities.Inventory{
			PlayerID: buyingReq.PlayerID,
			ItemID:   buyingReq.ItemID,
		})
	}

	return inventoryEntities

}

func (s *itemShopServiceImpl) checkPlayerItemQuantity(playerID string, itemID uint64, quantity uint) error {
	inventoryCount := s.inventoryRepository.PlayerItemCounting(playerID, itemID)

	if int(inventoryCount) < int(quantity) {
		log.Printf("Player %s has not enough item with id: %d", playerID, itemID)
		return &_itemShopException.NotEnoughItemException{ItemID: itemID}
	}

	return nil
}

func (s *itemShopServiceImpl) checkPlayerBalance(playerID string, amount int64) error {
	balanceDto, err := s.playerBalancingRepository.Showing(playerID)
	if err != nil {
		return err
	}

	if balanceDto.Balance < amount {
		log.Printf("Player %s has not enough balance", playerID)
		return &_itemShopException.NotEnoughBalanceException{}
	}

	return nil
}

func (s *itemShopServiceImpl) calculateTotalPrice(item *_itemShopModel.Item, quantity uint) int64 {
	// In a real world scenario, this would be a more complex calculation
	return int64(item.Price) * int64(quantity)
}

func (s *itemShopServiceImpl) totalPageCalculation(totalItems, size int64) int64 {
	totalPage := totalItems / size

	if totalItems%size != 0 {
		totalPage++
	}

	return totalPage
}

func (s *itemShopServiceImpl) buildItemResultsResponse(itemEntityList []*entities.Item, page, totalPage int64) *_itemShopModel.ItemResult {
	items := make([]*_itemShopModel.Item, 0)

	for _, itemEntity := range itemEntityList {
		items = append(items, itemEntity.ToItemModel())
	}

	return &_itemShopModel.ItemResult{
		Items: items,
		Paginate: _itemShopModel.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}
