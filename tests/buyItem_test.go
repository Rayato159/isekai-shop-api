package tests

import (
	_historyOfPurchasingEntity "github.com/Rayato159/isekai-shop-api/domains/historyOfPurchasing/entity"
	_historyOfPurchasingRepository "github.com/Rayato159/isekai-shop-api/domains/historyOfPurchasing/repository"
	_itemEntity "github.com/Rayato159/isekai-shop-api/domains/item/entity"
	_itemRepository "github.com/Rayato159/isekai-shop-api/domains/item/repository"
	_paymentEntity "github.com/Rayato159/isekai-shop-api/domains/payment/entity"
	_paymentException "github.com/Rayato159/isekai-shop-api/domains/payment/exception"
	_paymentModel "github.com/Rayato159/isekai-shop-api/domains/payment/model"
	_paymentRepository "github.com/Rayato159/isekai-shop-api/domains/payment/repository"
	_paymentService "github.com/Rayato159/isekai-shop-api/domains/payment/service"
	_playerEntity "github.com/Rayato159/isekai-shop-api/domains/player/entity"
	_playerSource "github.com/Rayato159/isekai-shop-api/domains/player/repository"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestItemBuyingSuccess(t *testing.T) {
	itemRepositoryMock := new(_itemRepository.ItemRepositoryMock)
	historyOfPurchasingRepositoryMock := new(_historyOfPurchasingRepository.HistoryOfPurchasingRepositoryMock)
	paymentRepositoryMock := new(_paymentRepository.PaymentRepositoryMock)
	inventoryRepositoryMock := new(_playerSource.InventoryRepositoryMock)

	paymentService := _paymentService.NewPaymentServiceImpl(
		paymentRepositoryMock,
		itemRepositoryMock,
		historyOfPurchasingRepositoryMock,
		inventoryRepositoryMock,
	)

	itemRepositoryMock.On("FindItemByID", uint64(1)).Return(&_itemEntity.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	paymentRepositoryMock.On("PlayerBalanceShowing", "P001").Return(&_paymentEntity.PlayerBalanceDto{
		PlayerID: "P001",
		Balance:  5000,
	}, nil)

	historyOfPurchasingRepositoryMock.On("HistoryOfPurchasingRecording", &_historyOfPurchasingEntity.HistoryOfPurchasing{
		PlayerID:        "P001",
		ItemID:          1,
		ItemName:        "Sword of Tester",
		ItemDescription: "A sword that can be used to test the enemy's defense",
		ItemPicture:     "https://www.google.com/sword-of-tester.jpg",
		ItemPrice:       1000,
		Quantity:        3,
	}).Return(&_historyOfPurchasingEntity.HistoryOfPurchasing{
		PlayerID:        "P001",
		ItemID:          1,
		ItemName:        "Sword of Tester",
		ItemDescription: "A sword that can be used to test the enemy's defense",
		ItemPicture:     "https://www.google.com/sword-of-tester.jpg",
		ItemPrice:       1000,
		Quantity:        3,
	}, nil)

	inventoryRepositoryMock.On("InventoryFilling", []*_playerEntity.Inventory{
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
	}).Return([]*_playerEntity.Inventory{
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

	paymentRepositoryMock.On("PaymentRecording", &_paymentEntity.Payment{
		PlayerID: "P001",
		Amount:   -3000,
	}).Return(&_paymentEntity.Payment{
		PlayerID: "P001",
		Amount:   -3000,
	}, nil)

	type args struct {
		in       *_paymentModel.ItemBuyingReq
		expected *_paymentModel.Payment
	}

	cases := []args{
		{
			in: &_paymentModel.ItemBuyingReq{
				PlayerID: "P001",
				ItemID:   1,
				Quantity: 3,
			},
			expected: &_paymentModel.Payment{
				PlayerID: "P001",
				Amount:   -3000,
			},
		},
	}

	for _, c := range cases {
		result, err := paymentService.ItemBuying(c.in)
		assert.NoError(t, err)
		assert.EqualValues(t, c.expected, result)
	}
}

func TestItemBuyingFail(t *testing.T) {
	itemRepositoryMock := new(_itemRepository.ItemRepositoryMock)
	historyOfPurchasingRepositoryMock := new(_historyOfPurchasingRepository.HistoryOfPurchasingRepositoryMock)
	inventoryRepositoryMock := new(_playerSource.InventoryRepositoryMock)
	paymentRepositoryMock := new(_paymentRepository.PaymentRepositoryMock)

	paymentService := _paymentService.NewPaymentServiceImpl(
		paymentRepositoryMock,
		itemRepositoryMock,
		historyOfPurchasingRepositoryMock,
		inventoryRepositoryMock,
	)

	itemRepositoryMock.On("FindItemByID", uint64(1)).Return(&_itemEntity.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	paymentRepositoryMock.On("PlayerBalanceShowing", "P001").Return(&_paymentEntity.PlayerBalanceDto{
		PlayerID: "P001",
		Balance:  2000,
	}, nil)

	type args struct {
		in       *_paymentModel.ItemBuyingReq
		expected error
	}

	cases := []args{
		{
			in: &_paymentModel.ItemBuyingReq{
				PlayerID: "P001",
				ItemID:   1,
				Quantity: 3,
			},
			expected: &_paymentException.NotEnoughBalanceException{},
		},
	}

	for _, c := range cases {
		result, err := paymentService.ItemBuying(c.in)
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.EqualValues(t, c.expected, err)
	}
}
