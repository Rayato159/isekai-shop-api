package service

import (
	_inventoryEntity "github.com/Rayato159/isekai-shop-api/modules/inventory/entity"
	_inventoryRepository "github.com/Rayato159/isekai-shop-api/modules/inventory/repository"
	_itemModel "github.com/Rayato159/isekai-shop-api/modules/item/model"
	_itemRepository "github.com/Rayato159/isekai-shop-api/modules/item/repository"
	_orderEntity "github.com/Rayato159/isekai-shop-api/modules/order/entity"
	_orderRepository "github.com/Rayato159/isekai-shop-api/modules/order/repository"
	_paymentEntity "github.com/Rayato159/isekai-shop-api/modules/payment/entity"
	_paymentModel "github.com/Rayato159/isekai-shop-api/modules/payment/model"
	_paymentRepository "github.com/Rayato159/isekai-shop-api/modules/payment/repository"
	"github.com/labstack/echo/v4"
)

type paymentServiceImpl struct {
	paymentRepository   _paymentRepository.PaymentRepository
	itemRepository      _itemRepository.ItemRepository
	orderRepository     _orderRepository.OrderRepository
	inventoryRepository _inventoryRepository.InventoryRepository
	logger              echo.Logger
}

func NewPaymentServiceImpl(
	paymentRepository _paymentRepository.PaymentRepository,
	itemRepository _itemRepository.ItemRepository,
	orderRepository _orderRepository.OrderRepository,
	inventoryRepository _inventoryRepository.InventoryRepository,
	logger echo.Logger,
) PaymentService {
	return &paymentServiceImpl{
		paymentRepository:   paymentRepository,
		itemRepository:      itemRepository,
		orderRepository:     orderRepository,
		inventoryRepository: inventoryRepository,
		logger:              logger,
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
	s.logger.Infof("Inserted order: %s", insertedOrder.ID)

	inventoryEntities := s.groupInventoryEntities(buyItemReq)

	inventory, err := s.inventoryRepository.InsertInventoryInBluk(inventoryEntities)
	if err != nil {
		return nil, err
	}
	s.logger.Infof("Inserted inventories: %d", len(inventory))

	insertedPayment, err := s.paymentRepository.InsertPayment(&_paymentEntity.Payment{
		PlayerID: buyItemReq.PlayerID,
		Amount:   -totalPrice,
	})
	if err != nil {
		return nil, err
	}
	s.logger.Infof("Payment entity: %s", insertedPayment.ID)

	return insertedPayment.ToPaymentModel(), nil
}

// 1. Search item by ID
// 2. Calculate total price
// 3. Check if player has enough quantity
// 4. Create order
// 5. Create payment
// 6. Delete item into player inventory
func (s *paymentServiceImpl) SellItem(sellItemReq *_paymentModel.SellItemReq) (*_paymentModel.Payment, error) {
	return nil, nil
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

func (s *paymentServiceImpl) calculateTotalPrice(item *_itemModel.Item, quantity uint) int64 {
	// In a real world scenario, this would be a more complex calculation
	return int64(item.Price) * int64(quantity)
}
