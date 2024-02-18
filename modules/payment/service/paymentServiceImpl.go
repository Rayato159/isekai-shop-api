package service

import (
	"log"

	_inventoryEntity "github.com/Rayato159/isekai-shop-api/modules/inventory/entity"
	_inventoryRepository "github.com/Rayato159/isekai-shop-api/modules/inventory/repository"
	_itemModel "github.com/Rayato159/isekai-shop-api/modules/item/model"
	_itemRepository "github.com/Rayato159/isekai-shop-api/modules/item/repository"
	_orderEntity "github.com/Rayato159/isekai-shop-api/modules/order/entity"
	_orderRepository "github.com/Rayato159/isekai-shop-api/modules/order/repository"
	_paymentEntity "github.com/Rayato159/isekai-shop-api/modules/payment/entity"
	_paymentException "github.com/Rayato159/isekai-shop-api/modules/payment/exception"
	_paymentModel "github.com/Rayato159/isekai-shop-api/modules/payment/model"
	_paymentRepository "github.com/Rayato159/isekai-shop-api/modules/payment/repository"
)

type paymentServiceImpl struct {
	paymentRepository   _paymentRepository.PaymentRepository
	itemRepository      _itemRepository.ItemRepository
	orderRepository     _orderRepository.OrderRepository
	inventoryRepository _inventoryRepository.InventoryRepository
}

func NewPaymentServiceImpl(
	paymentRepository _paymentRepository.PaymentRepository,
	itemRepository _itemRepository.ItemRepository,
	orderRepository _orderRepository.OrderRepository,
	inventoryRepository _inventoryRepository.InventoryRepository,
) PaymentService {
	return &paymentServiceImpl{
		paymentRepository:   paymentRepository,
		itemRepository:      itemRepository,
		orderRepository:     orderRepository,
		inventoryRepository: inventoryRepository,
	}
}

func (s *paymentServiceImpl) TopUp(topUpReq *_paymentModel.TopUpReq) (*_paymentModel.Payment, error) {
	paymentEntity := &_paymentEntity.Payment{
		PlayerID: topUpReq.PlayerID,
		Amount:   topUpReq.Amount,
	}

	insertedPayment, err := s.paymentRepository.InsertPayment(paymentEntity)
	if err != nil {
		return nil, err
	}

	return insertedPayment.ToPaymentModel(), nil
}

func (s *paymentServiceImpl) CalculatePlayerBalance(playerID string) *_paymentModel.PlayerBalance {
	balanceDto, err := s.paymentRepository.CalculatePlayerBalance(playerID)
	if err != nil {
		return &_paymentModel.PlayerBalance{
			PlayerID: playerID,
			Balance:  0,
		}
	}

	return balanceDto.ToPlayerBalanceModel()
}

// 1. Search item by ID
// 2. Calculate total price
// 3. Check if player has enough balance
// 4. Create order
// 5. Create payment
// 6. Add item into player inventory
func (s *paymentServiceImpl) BuyItem(buyItemReq *_paymentModel.BuyItemReq) (*_paymentModel.Payment, error) {
	itemEntity, err := s.itemRepository.FindItemByID(buyItemReq.ItemID)
	if err != nil {
		return nil, err
	}

	totalPrice := s.calculateTotalPrice(itemEntity.ToItemModel(), buyItemReq.Quantity)

	if err := s.checkPlayerBalance(buyItemReq.PlayerID, totalPrice); err != nil {
		return nil, err
	}

	insertedOrder, err := s.orderRepository.InsertOrder(&_orderEntity.Order{
		PlayerID:        buyItemReq.PlayerID,
		ItemID:          itemEntity.ID,
		ItemName:        itemEntity.Name,
		ItemDescription: itemEntity.Description,
		ItemPrice:       itemEntity.Price,
		ItemPicture:     itemEntity.Picture,
		Quantity:        buyItemReq.Quantity,
		TotalPrice:      -totalPrice,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Inserted order: %d", insertedOrder.ID)

	inventoryEntities := s.groupInventoryEntities(buyItemReq)

	insertedPayment, err := s.paymentRepository.InsertPayment(&_paymentEntity.Payment{
		PlayerID: buyItemReq.PlayerID,
		Amount:   -totalPrice,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Payment entity: %d", insertedPayment.ID)

	inventory, err := s.inventoryRepository.InsertInventoryInBluk(inventoryEntities)
	if err != nil {
		return nil, err
	}
	log.Printf("Inserted inventories: %d", len(inventory))

	return insertedPayment.ToPaymentModel(), nil
}

// 1. Check if player has enough quantity
// 2. Search item by ID
// 3. Calculate total price
// 4. Create order
// 5. Create payment
// 6. Delete item into player inventory
func (s *paymentServiceImpl) SellItem(sellItemReq *_paymentModel.SellItemReq) (*_paymentModel.Payment, error) {
	if err := s.checkPlayerItemQuantity(
		sellItemReq.PlayerID,
		sellItemReq.ItemID,
		sellItemReq.Quantity,
	); err != nil {
		return nil, err
	}

	itemEntity, err := s.itemRepository.FindItemByID(sellItemReq.ItemID)
	if err != nil {
		return nil, err
	}

	totalPrice := s.calculateTotalPrice(itemEntity.ToItemModel(), sellItemReq.Quantity)
	totalPrice = totalPrice / 2

	insertedOrder, err := s.orderRepository.InsertOrder(&_orderEntity.Order{
		PlayerID:        sellItemReq.PlayerID,
		ItemID:          itemEntity.ID,
		ItemName:        itemEntity.Name,
		ItemDescription: itemEntity.Description,
		ItemPrice:       itemEntity.Price,
		ItemPicture:     itemEntity.Picture,
		Quantity:        sellItemReq.Quantity,
		TotalPrice:      totalPrice,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Inserted order: %d", insertedOrder.ID)

	insertedPayment, err := s.paymentRepository.InsertPayment(&_paymentEntity.Payment{
		PlayerID: sellItemReq.PlayerID,
		Amount:   totalPrice,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Payment entity: %d", insertedPayment.ID)

	if err := s.inventoryRepository.DeleteItemByLimit(
		sellItemReq.PlayerID,
		sellItemReq.ItemID,
		int(sellItemReq.Quantity),
	); err != nil {
		return nil, err
	}
	log.Printf("Deleted inventories for %d records", sellItemReq.Quantity)

	return insertedPayment.ToPaymentModel(), nil
}

func (s *paymentServiceImpl) groupInventoryEntities(buyItemReq *_paymentModel.BuyItemReq) []*_inventoryEntity.Inventory {
	inventoryEntities := make([]*_inventoryEntity.Inventory, 0)

	for i := 0; i < int(buyItemReq.Quantity); i++ {
		inventoryEntities = append(inventoryEntities, &_inventoryEntity.Inventory{
			PlayerID: buyItemReq.PlayerID,
			ItemID:   buyItemReq.ItemID,
		})
	}

	return inventoryEntities

}

func (s *paymentServiceImpl) checkPlayerItemQuantity(playerID string, itemID uint64, quantity uint) error {
	inventoryCount := s.inventoryRepository.CountPlayerItem(playerID, itemID)

	if int(inventoryCount) < int(quantity) {
		log.Printf("Player %s has not enough item with id: %d", playerID, itemID)
		return &_paymentException.NotEnoughItemException{ItemID: itemID}
	}

	return nil
}

func (s *paymentServiceImpl) checkPlayerBalance(playerID string, amount int64) error {
	balanceDto, err := s.paymentRepository.CalculatePlayerBalance(playerID)
	if err != nil {
		return err
	}

	if balanceDto.Balance < amount {
		log.Printf("Player %s has not enough balance", playerID)
		return &_paymentException.NotEnoughBalanceException{}
	}

	return nil
}

func (s *paymentServiceImpl) calculateTotalPrice(item *_itemModel.Item, quantity uint) int64 {
	// In a real world scenario, this would be a more complex calculation
	return int64(item.Price) * int64(quantity)
}
