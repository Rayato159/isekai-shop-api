package tests

import (
	_inventoryRepository "github.com/Rayato159/isekai-shop-api/domains/inventory/repository"
	_itemShopException "github.com/Rayato159/isekai-shop-api/domains/itemShop/exception"
	_itemShopModel "github.com/Rayato159/isekai-shop-api/domains/itemShop/model"
	_itemShopRepository "github.com/Rayato159/isekai-shop-api/domains/itemShop/repository"
	_itemShopService "github.com/Rayato159/isekai-shop-api/domains/itemShop/service"
	_playerCoinModel "github.com/Rayato159/isekai-shop-api/domains/playerCoin/model"
	_playerCoinRepository "github.com/Rayato159/isekai-shop-api/domains/playerCoin/repository"
	entities "github.com/Rayato159/isekai-shop-api/entities"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestItemBuyingSuccess(t *testing.T) {
	itemShopRepositoryMock := new(_itemShopRepository.ItemShopRepositoryMock)
	playerCoinRepositoryMock := new(_playerCoinRepository.CoinRepositoryMock)
	inventoryRepositoryMock := new(_inventoryRepository.InventoryRepositoryMock)

	itemShopService := _itemShopService.NewItemShopServiceImpl(
		itemShopRepositoryMock,
		playerCoinRepositoryMock,
		inventoryRepositoryMock,
	)

	itemShopRepositoryMock.On("FindByID", uint64(1)).Return(&entities.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	playerCoinRepositoryMock.On("Showing", "P001").Return(&entities.PlayerBalanceShowingDto{
		PlayerID: "P001",
		Balance:  5000,
	}, nil)

	itemShopRepositoryMock.On("PurchasingHistoryRecording", &entities.PurchasingHistory{
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

	playerCoinRepositoryMock.On("Recording", &entities.PlayerCoin{
		PlayerID: "P001",
		Amount:   -3000,
	}).Return(&entities.PlayerCoin{
		PlayerID: "P001",
		Amount:   -3000,
	}, nil)

	type args struct {
		in       *_itemShopModel.BuyingReq
		expected *_playerCoinModel.PlayerCoin
	}

	cases := []args{
		{
			in: &_itemShopModel.BuyingReq{
				PlayerID: "P001",
				ItemID:   1,
				Quantity: 3,
			},
			expected: &_playerCoinModel.PlayerCoin{
				PlayerID: "P001",
				Amount:   -3000,
			},
		},
	}

	for _, c := range cases {
		result, err := itemShopService.Buying(c.in)
		assert.NoError(t, err)
		assert.EqualValues(t, c.expected, result)
	}
}

func TestItemBuyingFail(t *testing.T) {
	itemShopRepositoryMock := new(_itemShopRepository.ItemShopRepositoryMock)
	inventoryRepositoryMock := new(_inventoryRepository.InventoryRepositoryMock)
	playerCoinRepositoryMock := new(_playerCoinRepository.CoinRepositoryMock)

	itemShopService := _itemShopService.NewItemShopServiceImpl(
		itemShopRepositoryMock,
		playerCoinRepositoryMock,
		inventoryRepositoryMock,
	)

	itemShopRepositoryMock.On("FindByID", uint64(1)).Return(&entities.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	playerCoinRepositoryMock.On("Showing", "P001").Return(&entities.PlayerBalanceShowingDto{
		PlayerID: "P001",
		Balance:  2000,
	}, nil)

	type args struct {
		in       *_itemShopModel.BuyingReq
		expected error
	}

	cases := []args{
		{
			in: &_itemShopModel.BuyingReq{
				PlayerID: "P001",
				ItemID:   1,
				Quantity: 3,
			},
			expected: &_itemShopException.NotEnoughBalanceException{},
		},
	}

	for _, c := range cases {
		result, err := itemShopService.Buying(c.in)
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.EqualValues(t, c.expected, err)
	}
}
