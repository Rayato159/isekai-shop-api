package tests

import (
	_balancingModel "github.com/Rayato159/isekai-shop-api/domains/balancing/model"
	_balancingRepository "github.com/Rayato159/isekai-shop-api/domains/balancing/repository"
	_inventoryRepository "github.com/Rayato159/isekai-shop-api/domains/inventory/repository"
	_itemGettingRepository "github.com/Rayato159/isekai-shop-api/domains/itemGetting/repository"
	_puchasingException "github.com/Rayato159/isekai-shop-api/domains/purchasing/exception"
	_puchasingModel "github.com/Rayato159/isekai-shop-api/domains/purchasing/model"
	_purchasingRepository "github.com/Rayato159/isekai-shop-api/domains/purchasing/repository"
	_purchasingService "github.com/Rayato159/isekai-shop-api/domains/purchasing/service"
	entities "github.com/Rayato159/isekai-shop-api/entities"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestItemSellingSuccess(t *testing.T) {
	itemRepositoryMock := new(_itemGettingRepository.ItemGettingRepositoryMock)
	purchasingRepositoryMock := new(_purchasingRepository.PurchasingRepositoryMock)
	balancingRepositoryMock := new(_balancingRepository.BalancingRepositoryMock)
	inventoryRepositoryMock := new(_inventoryRepository.InventoryRepositoryMock)

	purchasingService := _purchasingService.NewPurchasingServiceImpl(
		balancingRepositoryMock,
		itemRepositoryMock,
		purchasingRepositoryMock,
		inventoryRepositoryMock,
	)

	inventoryRepositoryMock.On("PlayerItemCounting", "P001", uint64(1)).Return(int64(3), nil)

	itemRepositoryMock.On("FindByID", uint64(1)).Return(&entities.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	purchasingRepositoryMock.On("PurchasingHistoryRecording", &entities.PurchasingHistory{
		PlayerID:        "P001",
		ItemID:          1,
		ItemName:        "Sword of Tester",
		ItemDescription: "A sword that can be used to test the enemy's defense",
		ItemPicture:     "https://www.google.com/sword-of-tester.jpg",
		ItemPrice:       1000,
		Quantity:        3,
	}).Return(&entities.PurchasingHistory{
		PlayerID:        "P001",
		ItemID:          1,
		ItemName:        "Sword of Tester",
		ItemDescription: "A sword that can be used to test the enemy's defense",
		ItemPicture:     "https://www.google.com/sword-of-tester.jpg",
		ItemPrice:       1000,
		Quantity:        3,
	}, nil)

	balancingRepositoryMock.On("PlayerBalanceRecording", &entities.Balancing{
		PlayerID: "P001",
		Amount:   1500,
	}).Return(&entities.Balancing{
		PlayerID: "P001",
		Amount:   1500,
	}, nil)

	inventoryRepositoryMock.On("DeletePlayerItemByLimit", "P001", uint64(1), 3).Return(nil)

	type args struct {
		in       *_puchasingModel.ItemSellingReq
		expected *_balancingModel.Balancing
	}

	cases := []args{
		{
			&_puchasingModel.ItemSellingReq{
				PlayerID: "P001",
				ItemID:   1,
				Quantity: 3,
			},
			&_balancingModel.Balancing{
				PlayerID: "P001",
				Amount:   1500,
			},
		},
	}

	for _, c := range cases {
		result, err := purchasingService.ItemSelling(c.in)
		assert.NoError(t, err)
		assert.Equal(t, c.expected, result)
	}
}

func TestItemSellingFailed(t *testing.T) {
	itemRepositoryMock := new(_itemGettingRepository.ItemGettingRepositoryMock)
	purchasingRepositoryMock := new(_purchasingRepository.PurchasingRepositoryMock)
	balancingRepositoryMock := new(_balancingRepository.BalancingRepositoryMock)
	inventoryRepositoryMock := new(_inventoryRepository.InventoryRepositoryMock)

	purchasingService := _purchasingService.NewPurchasingServiceImpl(
		balancingRepositoryMock,
		itemRepositoryMock,
		purchasingRepositoryMock,
		inventoryRepositoryMock,
	)

	inventoryRepositoryMock.On("PlayerItemCounting", "P001", uint64(1)).Return(int64(2), nil)

	type args struct {
		in       *_puchasingModel.ItemSellingReq
		expected error
	}

	cases := []args{
		{
			&_puchasingModel.ItemSellingReq{
				PlayerID: "P001",
				ItemID:   1,
				Quantity: 3,
			},
			&_puchasingException.NotEnoughItemException{ItemID: 1},
		},
	}

	for _, c := range cases {
		result, err := purchasingService.ItemSelling(c.in)
		assert.EqualValues(t, c.expected, err)
		assert.Nil(t, result)
	}
}
