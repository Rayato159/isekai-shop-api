package service

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"
	_inventoryRepository "github.com/Rayato159/isekai-shop-api/pkg/inventory/repository"
	_itemShopException "github.com/Rayato159/isekai-shop-api/pkg/itemShop/exception"
	_itemShopModel "github.com/Rayato159/isekai-shop-api/pkg/itemShop/model"
	_itemShopRepository "github.com/Rayato159/isekai-shop-api/pkg/itemShop/repository"
	_playerCoinModel "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/model"
	_playerCoinRepository "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/repository"
	"github.com/labstack/echo/v4"
)

type itemShopServiceImpl struct {
	itemShopRepository   _itemShopRepository.ItemShopRepository
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository
	inventoryRepository  _inventoryRepository.InventoryRepository
	logger               echo.Logger
}

func NewItemShopServiceImpl(
	itemShopRepository _itemShopRepository.ItemShopRepository,
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository,
	inventoryRepository _inventoryRepository.InventoryRepository,
	logger echo.Logger,
) ItemShopService {
	return &itemShopServiceImpl{
		itemShopRepository,
		playerCoinRepository,
		inventoryRepository,
		logger,
	}
}

func (s *itemShopServiceImpl) Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error) {
	itemEntityList, err := s.itemShopRepository.Listing(itemFilter)
	if err != nil {
		return nil, err
	}

	totalItems, err := s.itemShopRepository.Counting(itemFilter)
	if err != nil {
		return nil, err
	}

	size := itemFilter.Paginate.Size
	page := itemFilter.Paginate.Page
	totalPage := s.totalPageCalculation(totalItems, size)

	result := s.toItemResultsResponse(itemEntityList, page, totalPage)

	return result, nil
}

// 1. Search item by ID
// 2. Calculate total price
// 3. Check if player has enough coin
// 4. Create itemShop
// 5. Create playerCoin
// 6. Add item into player inventory
func (s *itemShopServiceImpl) Buying(buyingReq *_itemShopModel.BuyingReq) (*_playerCoinModel.PlayerCoin, error) {
	itemEntity, err := s.itemShopRepository.FindByID(buyingReq.ItemID)
	if err != nil {
		return nil, err
	}

	totalPrice := s.calculateTotalPrice(itemEntity.ToItemModel(), buyingReq.Quantity)

	if err := s.playerCoinChecking(buyingReq.PlayerID, totalPrice); err != nil {
		return nil, err
	}

	tx := s.itemShopRepository.BeginTransaction()

	purchaseRecording, err := s.itemShopRepository.PurchaseHistoryRecording(&entities.PurchaseHistory{
		PlayerID:        buyingReq.PlayerID,
		ItemID:          itemEntity.ID,
		ItemName:        itemEntity.Name,
		ItemDescription: itemEntity.Description,
		ItemPrice:       itemEntity.Price,
		ItemPicture:     itemEntity.Picture,
		Quantity:        buyingReq.Quantity,
		IsBuying:        true,
	}, tx)
	if err != nil {
		s.itemShopRepository.RollbackTransaction(tx)
		return nil, err
	}
	s.logger.Infof("Purchase history recorded: %d", purchaseRecording.ID)

	coinRecording, err := s.playerCoinRepository.CoinAdding(&entities.PlayerCoin{
		PlayerID: buyingReq.PlayerID,
		Amount:   -totalPrice,
	}, tx)
	if err != nil {
		s.itemShopRepository.RollbackTransaction(tx)
		return nil, err
	}
	s.logger.Infof("Player coins reduced for: %d", totalPrice)

	inventoryRecording, err := s.inventoryRepository.Filling(
		buyingReq.PlayerID,
		buyingReq.ItemID,
		int(buyingReq.Quantity),
		tx,
	)
	if err != nil {
		s.itemShopRepository.RollbackTransaction(tx)
		return nil, err
	}
	s.logger.Infof("Items recorded into player inventory: %d", len(inventoryRecording))

	if err := s.itemShopRepository.CommitTransaction(tx); err != nil {
		s.itemShopRepository.RollbackTransaction(tx)
		return nil, err
	}

	return coinRecording.ToPlayerCoinModel(), nil
}

// 1. Check if player has enough quantity
// 2. Search item by ID
// 3. Calculate total price
// 4. Create itemShop
// 5. Create playerCoin
// 6. Delete item into player inventory
func (s *itemShopServiceImpl) Selling(sellingReq *_itemShopModel.SellingReq) (*_playerCoinModel.PlayerCoin, error) {
	itemEntity, err := s.itemShopRepository.FindByID(sellingReq.ItemID)
	if err != nil {
		return nil, err
	}

	totalPrice := s.calculateTotalPrice(itemEntity.ToItemModel(), sellingReq.Quantity)
	totalPrice = totalPrice / 2

	if err := s.playerItemChecking(
		sellingReq.PlayerID,
		sellingReq.ItemID,
		sellingReq.Quantity,
	); err != nil {
		return nil, err
	}

	tx := s.itemShopRepository.BeginTransaction()

	purchaseRecording, err := s.itemShopRepository.PurchaseHistoryRecording(&entities.PurchaseHistory{
		PlayerID:        sellingReq.PlayerID,
		ItemID:          itemEntity.ID,
		ItemName:        itemEntity.Name,
		ItemDescription: itemEntity.Description,
		ItemPrice:       itemEntity.Price,
		ItemPicture:     itemEntity.Picture,
		Quantity:        sellingReq.Quantity,
		IsBuying:        false,
	}, tx)
	if err != nil {
		s.itemShopRepository.RollbackTransaction(tx)
		return nil, err
	}
	s.logger.Infof("Puchase history recorded: %d", purchaseRecording.ID)

	coinRecording, err := s.playerCoinRepository.CoinAdding(&entities.PlayerCoin{
		PlayerID: sellingReq.PlayerID,
		Amount:   totalPrice,
	}, tx)
	if err != nil {
		s.itemShopRepository.RollbackTransaction(tx)
		return nil, err
	}
	s.logger.Infof("Coins added into player for: %d", coinRecording.Amount)

	if err := s.inventoryRepository.Removing(
		sellingReq.PlayerID,
		sellingReq.ItemID,
		int(sellingReq.Quantity),
		tx,
	); err != nil {
		s.itemShopRepository.RollbackTransaction(tx)
		return nil, err
	}
	s.logger.Infof("Deleted player item from player's inventory for %d records", sellingReq.Quantity)

	if err := s.itemShopRepository.CommitTransaction(tx); err != nil {
		s.itemShopRepository.RollbackTransaction(tx)
		return nil, err
	}

	return coinRecording.ToPlayerCoinModel(), nil
}

func (s *itemShopServiceImpl) playerItemChecking(playerID string, itemID uint64, quantity uint) error {
	inventoryCount := s.inventoryRepository.PlayerItemCounting(playerID, itemID)

	if int(inventoryCount) < int(quantity) {
		s.logger.Errorf("Player %s has not enough item with id: %d", playerID, itemID)
		return &_itemShopException.ItemNotEnough{ItemID: itemID}
	}

	return nil
}

func (s *itemShopServiceImpl) playerCoinChecking(playerID string, totalPrice int64) error {
	playerCoin, err := s.playerCoinRepository.Showing(playerID)
	if err != nil {
		return err
	}

	if playerCoin.Coin < totalPrice {
		s.logger.Errorf("Player %s has not enough coin", playerID)
		return &_itemShopException.CoinNotEnough{}
	}

	return nil
}

func (s *itemShopServiceImpl) calculateTotalPrice(item *_itemShopModel.Item, qty uint) int64 {
	// In a real world scenario, this would be a more complex calculation
	return int64(item.Price) * int64(qty)
}

func (s *itemShopServiceImpl) totalPageCalculation(totalItems, size int64) int64 {
	totalPage := totalItems / size

	if totalItems%size != 0 {
		totalPage++
	}

	return totalPage
}

func (s *itemShopServiceImpl) toItemResultsResponse(itemEntityList []*entities.Item, page, totalPage int64) *_itemShopModel.ItemResult {
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
