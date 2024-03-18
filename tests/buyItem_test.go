package tests

import (
	_balancingModel "github.com/Rayato159/isekai-shop-api/domains/balancing/model"
	_balancingRepository "github.com/Rayato159/isekai-shop-api/domains/balancing/repository"
	entities "github.com/Rayato159/isekai-shop-api/domains/entities"
	_itemRepository "github.com/Rayato159/isekai-shop-api/domains/item/repository"
	_playerSource "github.com/Rayato159/isekai-shop-api/domains/player/repository"
	_purchasingException "github.com/Rayato159/isekai-shop-api/domains/purchasing/exception"
	_purchasingModel "github.com/Rayato159/isekai-shop-api/domains/purchasing/model"
	_purchasingRepository "github.com/Rayato159/isekai-shop-api/domains/purchasing/repository"
	_purchasingService "github.com/Rayato159/isekai-shop-api/domains/purchasing/service"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestItemBuyingSuccess(t *testing.T) {
	itemRepositoryMock := new(_itemRepository.ItemRepositoryMock)
	purchasingRepositoryMock := new(_purchasingRepository.PurchasingRepositoryMock)
	balancingRepositoryMock := new(_balancingRepository.BalancingRepositoryMock)
	inventoryRepositoryMock := new(_playerSource.InventoryRepositoryMock)

	purchasingService := _purchasingService.NewPurchasingServiceImpl(
		balancingRepositoryMock,
		itemRepositoryMock,
		purchasingRepositoryMock,
		inventoryRepositoryMock,
	)

	itemRepositoryMock.On("FindItemByID", uint64(1)).Return(&entities.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	balancingRepositoryMock.On("PlayerBalanceShowing", "P001").Return(&entities.PlayerBalanceDto{
		PlayerID: "P001",
		Balance:  5000,
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

	inventoryRepositoryMock.On("InventoryFilling", []*entities.Inventory{
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
	}).Return([]*entities.Inventory{
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

	balancingRepositoryMock.On("BalancingRecording", &entities.Balancing{
		PlayerID: "P001",
		Amount:   -3000,
	}).Return(&entities.Balancing{
		PlayerID: "P001",
		Amount:   -3000,
	}, nil)

	type args struct {
		in       *_purchasingModel.ItemBuyingReq
		expected *_balancingModel.Balancing
	}

	cases := []args{
		{
			in: &_purchasingModel.ItemBuyingReq{
				PlayerID: "P001",
				ItemID:   1,
				Quantity: 3,
			},
			expected: &_balancingModel.Balancing{
				PlayerID: "P001",
				Amount:   -3000,
			},
		},
	}

	for _, c := range cases {
		result, err := purchasingService.ItemBuying(c.in)
		assert.NoError(t, err)
		assert.EqualValues(t, c.expected, result)
	}
}

func TestItemBuyingFail(t *testing.T) {
	itemRepositoryMock := new(_itemRepository.ItemRepositoryMock)
	purchasingRepositoryMock := new(_purchasingRepository.PurchasingRepositoryMock)
	inventoryRepositoryMock := new(_playerSource.InventoryRepositoryMock)
	balancingRepositoryMock := new(_balancingRepository.BalancingRepositoryMock)

	purchasingService := _purchasingService.NewPurchasingServiceImpl(
		balancingRepositoryMock,
		itemRepositoryMock,
		purchasingRepositoryMock,
		inventoryRepositoryMock,
	)

	itemRepositoryMock.On("FindItemByID", uint64(1)).Return(&entities.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	balancingRepositoryMock.On("PlayerBalanceShowing", "P001").Return(&entities.PlayerBalanceDto{
		PlayerID: "P001",
		Balance:  2000,
	}, nil)

	type args struct {
		in       *_purchasingModel.ItemBuyingReq
		expected error
	}

	cases := []args{
		{
			in: &_purchasingModel.ItemBuyingReq{
				PlayerID: "P001",
				ItemID:   1,
				Quantity: 3,
			},
			expected: &_purchasingException.NotEnoughBalanceException{},
		},
	}

	for _, c := range cases {
		result, err := purchasingService.ItemBuying(c.in)
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.EqualValues(t, c.expected, err)
	}
}
