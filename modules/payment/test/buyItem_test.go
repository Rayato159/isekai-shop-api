package test

import (
	_inventoryEntity "github.com/Rayato159/isekai-shop-api/modules/inventory/entity"
	_inventoryRepository "github.com/Rayato159/isekai-shop-api/modules/inventory/repository"
	_itemEntity "github.com/Rayato159/isekai-shop-api/modules/item/entity"
	_itemRepository "github.com/Rayato159/isekai-shop-api/modules/item/repository"
	_orderEntity "github.com/Rayato159/isekai-shop-api/modules/order/entity"
	_orderRepository "github.com/Rayato159/isekai-shop-api/modules/order/repository"
	_paymentEntity "github.com/Rayato159/isekai-shop-api/modules/payment/entity"
	_paymentException "github.com/Rayato159/isekai-shop-api/modules/payment/exception"
	_paymentModel "github.com/Rayato159/isekai-shop-api/modules/payment/model"
	_paymentRepository "github.com/Rayato159/isekai-shop-api/modules/payment/repository"
	_paymentService "github.com/Rayato159/isekai-shop-api/modules/payment/service"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestBuyItemSuccess(t *testing.T) {
	itemRepositoryMock := new(_itemRepository.ItemRepositoryMock)
	orderRepositoryMock := new(_orderRepository.OrderRepositoryMock)
	inventoryRepositoryMock := new(_inventoryRepository.InventoryRepositoryMock)
	paymentRepositoryMock := new(_paymentRepository.PaymentRepositoryMock)

	itemRepositoryMock.On("FindItemByID", uint64(1)).Return(&_itemEntity.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	paymentRepositoryMock.On("CalculatePlayerBalance", "P001").Return(&_paymentEntity.PlayerBalanceDto{
		PlayerID: "P001",
		Balance:  5000,
	}, nil)

	orderRepositoryMock.On("InsertOrder", &_orderEntity.Order{
		PlayerID:        "P001",
		ItemID:          1,
		ItemName:        "Sword of Tester",
		ItemDescription: "A sword that can be used to test the enemy's defense",
		ItemPicture:     "https://www.google.com/sword-of-tester.jpg",
		ItemPrice:       1000,
		Quantity:        3,
		TotalPrice:      -3000,
	}).Return(&_orderEntity.Order{
		PlayerID:        "P001",
		ItemID:          1,
		ItemName:        "Sword of Tester",
		ItemDescription: "A sword that can be used to test the enemy's defense",
		ItemPicture:     "https://www.google.com/sword-of-tester.jpg",
		ItemPrice:       1000,
		Quantity:        3,
		TotalPrice:      -3000,
	}, nil)

	inventoryRepositoryMock.On("InsertInventoryInBluk", []*_inventoryEntity.Inventory{
		{
			PlayerID: "P001",
			ItemID:   1,
		},
		{
			PlayerID: "P001",
			ItemID:   1,
		},
		{
			PlayerID: "P001",
			ItemID:   1,
		},
	}).Return([]*_inventoryEntity.Inventory{
		{
			PlayerID: "P001",
			ItemID:   1,
		},
		{
			PlayerID: "P001",
			ItemID:   1,
		},
		{
			PlayerID: "P001",
			ItemID:   1,
		},
	}, nil)

	paymentRepositoryMock.On("InsertPayment", &_paymentEntity.Payment{
		PlayerID: "P001",
		Amount:   -3000,
	}).Return(&_paymentEntity.Payment{
		PlayerID: "P001",
		Amount:   -3000,
	}, nil)

	paymentService := _paymentService.NewPaymentServiceImpl(
		paymentRepositoryMock,
		itemRepositoryMock,
		orderRepositoryMock,
		inventoryRepositoryMock,
	)

	buyItemReq := &_paymentModel.BuyItemReq{
		PlayerID: "P001",
		ItemID:   1,
		Quantity: 3,
	}

	result, err := paymentService.BuyItem(buyItemReq)
	assert.NoError(t, err)

	assert.EqualValues(t, &_paymentModel.Payment{
		PlayerID: "P001",
		Amount:   -3000,
	}, result)
}

func TestBuyItemFail(t *testing.T) {
	itemRepositoryMock := new(_itemRepository.ItemRepositoryMock)
	orderRepositoryMock := new(_orderRepository.OrderRepositoryMock)
	inventoryRepositoryMock := new(_inventoryRepository.InventoryRepositoryMock)
	paymentRepositoryMock := new(_paymentRepository.PaymentRepositoryMock)

	paymentService := _paymentService.NewPaymentServiceImpl(
		paymentRepositoryMock,
		itemRepositoryMock,
		orderRepositoryMock,
		inventoryRepositoryMock,
	)

	itemRepositoryMock.On("FindItemByID", uint64(1)).Return(&_itemEntity.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	paymentRepositoryMock.On("CalculatePlayerBalance", "P001").Return(&_paymentEntity.PlayerBalanceDto{
		PlayerID: "P001",
		Balance:  2000,
	}, nil)

	buyItemReq := &_paymentModel.BuyItemReq{
		PlayerID: "P001",
		ItemID:   1,
		Quantity: 3,
	}

	_, err := paymentService.BuyItem(buyItemReq)
	assert.EqualValues(t, &_paymentException.NotEnoughBalanceException{}, err)
}
