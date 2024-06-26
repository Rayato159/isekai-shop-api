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

func TestItemSellingSuccess(t *testing.T) {
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

	inventoryRepositoryMock.On("PlayerItemCounting", "P001", uint64(1)).Return(int64(3), nil)

	itemShopRepositoryMock.On("FindByID", uint64(1)).Return(&entities.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	itemShopRepositoryMock.On("PurchaseHistoryRecording", &entities.PurchaseHistory{
		PlayerID:        "P001",
		ItemID:          1,
		ItemName:        "Sword of Tester",
		ItemDescription: "A sword that can be used to test the enemy's defense",
		ItemPicture:     "https://www.google.com/sword-of-tester.jpg",
		ItemPrice:       1000,
		Quantity:        3,
		IsBuying:        false,
	}, tx).Return(&entities.PurchaseHistory{
		PlayerID:        "P001",
		ItemID:          1,
		ItemName:        "Sword of Tester",
		ItemDescription: "A sword that can be used to test the enemy's defense",
		ItemPicture:     "https://www.google.com/sword-of-tester.jpg",
		ItemPrice:       1000,
		Quantity:        3,
		IsBuying:        false,
	}, nil)

	playerCoinRepositoryMock.On("CoinAdding", &entities.PlayerCoin{
		PlayerID: "P001",
		Amount:   1500,
	}, tx).Return(&entities.PlayerCoin{
		PlayerID: "P001",
		Amount:   1500,
	}, nil)

	inventoryRepositoryMock.On("Removing", "P001", uint64(1), 3, tx).Return(nil)

	type args struct {
		label    string
		in       *_itemShopModel.SellingReq
		expected *_playerCoinModel.PlayerCoin
	}

	cases := []args{
		{
			"Selling item success",
			&_itemShopModel.SellingReq{
				PlayerID: "P001",
				ItemID:   1,
				Quantity: 3,
			},
			&_playerCoinModel.PlayerCoin{
				PlayerID: "P001",
				Amount:   1500,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			result, err := itemShopService.Selling(c.in)
			assert.NoError(t, err)
			assert.Equal(t, c.expected, result)
		})
	}
}

func TestItemSellingFailed(t *testing.T) {
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

	inventoryRepositoryMock.On("PlayerItemCounting", "P001", uint64(1)).Return(int64(2), nil)

	itemShopRepositoryMock.On("FindByID", uint64(1)).Return(&entities.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	type args struct {
		label    string
		in       *_itemShopModel.SellingReq
		expected error
	}

	cases := []args{
		{
			"Selling item failed because the item is not enough",
			&_itemShopModel.SellingReq{
				PlayerID: "P001",
				ItemID:   1,
				Quantity: 3,
			},
			&_itemShop.ItemNotEnough{ItemID: 1},
		},
	}

	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			result, err := itemShopService.Selling(c.in)
			assert.Error(t, err)
			assert.Equal(t, c.expected, err)
			assert.Nil(t, result)
		})
	}
}
