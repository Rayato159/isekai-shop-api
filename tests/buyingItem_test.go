package tests

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"
	_inventoryRepository "github.com/Rayato159/isekai-shop-api/pkg/inventory/repository"
	_itemShop "github.com/Rayato159/isekai-shop-api/pkg/itemShop/exception"
	_itemShopModel "github.com/Rayato159/isekai-shop-api/pkg/itemShop/model"
	_itemShopRepository "github.com/Rayato159/isekai-shop-api/pkg/itemShop/repository"
	_itemShopService "github.com/Rayato159/isekai-shop-api/pkg/itemShop/service"
	_playerCoinModel "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/model"
	_playerCoinRepository "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/repository"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"testing"
)

func TestItemBuyingSuccess(t *testing.T) {
	itemShopRepositoryMock := new(_itemShopRepository.ItemShopRepositoryMock)
	playerCoinRepositoryMock := new(_playerCoinRepository.CoinRepositoryMock)
	inventoryRepositoryMock := new(_inventoryRepository.InventoryRepositoryMock)
	echoLogger := echo.New().Logger

	itemShopService := _itemShopService.NewItemShopServiceImpl(
		itemShopRepositoryMock,
		playerCoinRepositoryMock,
		inventoryRepositoryMock,
		echoLogger,
	)

	tx := &gorm.DB{}
	itemShopRepositoryMock.On("BeginTransaction").Return(tx)
	itemShopRepositoryMock.On("CommitTransaction", tx).Return(nil)
	itemShopRepositoryMock.On("RollbackTransaction", tx).Return(nil)

	itemShopRepositoryMock.On("FindByID", uint64(1)).Return(&entities.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	playerCoinRepositoryMock.On("Showing", "P001").Return(&_playerCoinModel.PlayerCoinShowing{
		PlayerID: "P001",
		Coin:     5000,
	}, nil)

	itemShopRepositoryMock.On("PurchaseHistoryRecording", &entities.PurchaseHistory{
		PlayerID:        "P001",
		ItemID:          1,
		ItemName:        "Sword of Tester",
		ItemDescription: "A sword that can be used to test the enemy's defense",
		ItemPicture:     "https://www.google.com/sword-of-tester.jpg",
		ItemPrice:       1000,
		Quantity:        3,
		IsBuying:        true,
	}, tx).Return(&entities.PurchaseHistory{
		PlayerID:        "P001",
		ItemID:          1,
		ItemName:        "Sword of Tester",
		ItemDescription: "A sword that can be used to test the enemy's defense",
		ItemPicture:     "https://www.google.com/sword-of-tester.jpg",
		ItemPrice:       1000,
		Quantity:        3,
		IsBuying:        true,
	}, nil)

	inventoryRepositoryMock.On("Filling", "P001", uint64(1), int(3), tx).Return([]*entities.Inventory{
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

	playerCoinRepositoryMock.On("CoinAdding", &entities.PlayerCoin{
		PlayerID: "P001",
		Amount:   -3000,
	}, tx).Return(&entities.PlayerCoin{
		PlayerID: "P001",
		Amount:   -3000,
	}, nil)

	type args struct {
		label    string
		in       *_itemShopModel.BuyingReq
		expected *_playerCoinModel.PlayerCoin
	}

	cases := []args{
		{
			label: "Success buying item",
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
		t.Run(c.label, func(t *testing.T) {
			result, err := itemShopService.Buying(c.in)
			assert.NoError(t, err)
			assert.EqualValues(t, c.expected, result)
		})
	}
}

func TestItemBuyingFail(t *testing.T) {
	itemShopRepositoryMock := new(_itemShopRepository.ItemShopRepositoryMock)
	inventoryRepositoryMock := new(_inventoryRepository.InventoryRepositoryMock)
	playerCoinRepositoryMock := new(_playerCoinRepository.CoinRepositoryMock)
	echoLogger := echo.New().Logger

	itemShopService := _itemShopService.NewItemShopServiceImpl(
		itemShopRepositoryMock,
		playerCoinRepositoryMock,
		inventoryRepositoryMock,
		echoLogger,
	)

	tx := &gorm.DB{}
	itemShopRepositoryMock.On("BeginTransaction").Return(tx)
	itemShopRepositoryMock.On("CommitTransaction", tx).Return(nil)
	itemShopRepositoryMock.On("RollbackTransaction", tx).Return(nil)

	itemShopRepositoryMock.On("FindByID", uint64(1)).Return(&entities.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	playerCoinRepositoryMock.On("Showing", "P001").Return(&_playerCoinModel.PlayerCoinShowing{
		PlayerID: "P001",
		Coin:     2000,
	}, nil)

	type args struct {
		label    string
		in       *_itemShopModel.BuyingReq
		expected error
	}

	cases := []args{
		{
			label: "Test failed to find item 1",
			in: &_itemShopModel.BuyingReq{
				PlayerID: "P001",
				ItemID:   1,
				Quantity: 3,
			},
			expected: &_itemShop.CoinNotEnough{},
		},
	}

	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			result, err := itemShopService.Buying(c.in)
			assert.Nil(t, result)
			assert.Error(t, err)
			assert.EqualValues(t, c.expected, err)
		})
	}
}
