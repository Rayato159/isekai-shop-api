package repository

import (
	"github.com/Rayato159/isekai-shop-api/entities"
	"gorm.io/gorm"

	_playerCoinModel "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/model"

	"github.com/stretchr/testify/mock"
)

type CoinRepositoryMock struct {
	mock.Mock
}

func (m *CoinRepositoryMock) CoinAdding(playerCoinEntity *entities.PlayerCoin, tx *gorm.DB) (*entities.PlayerCoin, error) {
	args := m.Called(playerCoinEntity, tx)
	return args.Get(0).(*entities.PlayerCoin), args.Error(1)
}

func (m *CoinRepositoryMock) Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error) {
	args := m.Called(playerID)
	return args.Get(0).(*_playerCoinModel.PlayerCoinShowing), args.Error(1)
}
