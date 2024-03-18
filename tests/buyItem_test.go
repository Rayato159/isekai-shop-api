package tests

import (
	_inventoryRepository "github.com/Rayato159/isekai-shop-api/domains/inventory/repository"
	_itemGettingRepository "github.com/Rayato159/isekai-shop-api/domains/itemGetting/repository"
	_playerBalancingModel "github.com/Rayato159/isekai-shop-api/domains/playerBalancing/model"
	_playerBalancingRepository "github.com/Rayato159/isekai-shop-api/domains/playerBalancing/repository"
	_purchasingException "github.com/Rayato159/isekai-shop-api/domains/purchasing/exception"
	_purchasingModel "github.com/Rayato159/isekai-shop-api/domains/purchasing/model"
	_purchasingRepository "github.com/Rayato159/isekai-shop-api/domains/purchasing/repository"
	_purchasingService "github.com/Rayato159/isekai-shop-api/domains/purchasing/service"
	entities "github.com/Rayato159/isekai-shop-api/entities"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestItemBuyingSuccess(t *testing.T) {
	itemRepositoryMock := new(_itemGettingRepository.ItemGettingRepositoryMock)
	purchasingRepositoryMock := new(_purchasingRepository.PurchasingRepositoryMock)
	playerBalancingRepositoryMock := new(_playerBalancingRepository.BalancingRepositoryMock)
	inventoryRepositoryMock := new(_inventoryRepository.InventoryRepositoryMock)

	purchasingService := _purchasingService.NewPurchasingServiceImpl(
		playerBalancingRepositoryMock,
		itemRepositoryMock,
		purchasingRepositoryMock,
		inventoryRepositoryMock,
	)

	itemRepositoryMock.On("FindByID", uint64(1)).Return(&entities.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	playerBalancingRepositoryMock.On("Showing", "P001").Return(&entities.PlayerBalanceShowingDto{
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

	inventoryRepositoryMock.On("Filling", []*entities.Inventory{
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

	playerBalancingRepositoryMock.On("Recording", &entities.PlayerBalancing{
		PlayerID: "P001",
		Amount:   -3000,
	}).Return(&entities.PlayerBalancing{
		PlayerID: "P001",
		Amount:   -3000,
	}, nil)

	type args struct {
		in       *_purchasingModel.ItemBuyingReq
		expected *_playerBalancingModel.PlayerBalancing
	}

	cases := []args{
		{
			in: &_purchasingModel.ItemBuyingReq{
				PlayerID: "P001",
				ItemID:   1,
				Quantity: 3,
			},
			expected: &_playerBalancingModel.PlayerBalancing{
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
	itemGettingRepositoryMock := new(_itemGettingRepository.ItemGettingRepositoryMock)
	purchasingRepositoryMock := new(_purchasingRepository.PurchasingRepositoryMock)
	inventoryRepositoryMock := new(_inventoryRepository.InventoryRepositoryMock)
	playerBalancingRepositoryMock := new(_playerBalancingRepository.BalancingRepositoryMock)

	purchasingService := _purchasingService.NewPurchasingServiceImpl(
		playerBalancingRepositoryMock,
		itemGettingRepositoryMock,
		purchasingRepositoryMock,
		inventoryRepositoryMock,
	)

	itemGettingRepositoryMock.On("FindByID", uint64(1)).Return(&entities.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	playerBalancingRepositoryMock.On("Showing", "P001").Return(&entities.PlayerBalanceShowingDto{
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
