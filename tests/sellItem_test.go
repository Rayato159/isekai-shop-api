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
	_playerSource "github.com/Rayato159/isekai-shop-api/domains/player/repository"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestItemSellingSuccess(t *testing.T) {
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

	inventoryRepositoryMock.On("PlayerItemCounting", "P001", uint64(1)).Return(int64(3), nil)

	itemRepositoryMock.On("FindItemByID", uint64(1)).Return(&_itemEntity.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
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

	paymentRepositoryMock.On("PaymentRecording", &_paymentEntity.Payment{
		PlayerID: "P001",
		Amount:   1500,
	}).Return(&_paymentEntity.Payment{
		PlayerID: "P001",
		Amount:   1500,
	}, nil)

	inventoryRepositoryMock.On("DeleteItemByLimit", "P001", uint64(1), 3).Return(nil)

	type args struct {
		in       *_paymentModel.ItemSellingReq
		expected *_paymentModel.Payment
	}

	cases := []args{
		{
			&_paymentModel.ItemSellingReq{
				PlayerID: "P001",
				ItemID:   1,
				Quantity: 3,
			},
			&_paymentModel.Payment{
				PlayerID: "P001",
				Amount:   1500,
			},
		},
	}

	for _, c := range cases {
		result, err := paymentService.ItemSelling(c.in)
		assert.NoError(t, err)
		assert.Equal(t, c.expected, result)
	}
}

func TestItemSellingFailed(t *testing.T) {
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

	inventoryRepositoryMock.On("PlayerItemCounting", "P001", uint64(1)).Return(int64(2), nil)

	type args struct {
		in       *_paymentModel.ItemSellingReq
		expected error
	}

	cases := []args{
		{
			&_paymentModel.ItemSellingReq{
				PlayerID: "P001",
				ItemID:   1,
				Quantity: 3,
			},
			&_paymentException.NotEnoughItemException{ItemID: 1},
		},
	}

	for _, c := range cases {
		result, err := paymentService.ItemSelling(c.in)
		assert.EqualValues(t, c.expected, err)
		assert.Nil(t, result)
	}
}
