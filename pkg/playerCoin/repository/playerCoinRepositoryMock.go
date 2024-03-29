package repository

import (
	"github.com/Rayato159/isekai-shop-api/entities"

	_playerCoinModel "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/model"

	"github.com/stretchr/testify/mock"
)

type CoinRepositoryMock struct {
	mock.Mock
}

func (m *CoinRepositoryMock) CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error) {
	args := m.Called(playerCoinEntity)
	return args.Get(0).(*entities.PlayerCoin), args.Error(1)
}

func (m *CoinRepositoryMock) ReverseCoinAdding(playerCoinEntity *entities.PlayerCoin) error {
	args := m.Called(playerCoinEntity)
	return args.Error(0)
}

func (m *CoinRepositoryMock) Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error) {
	args := m.Called(playerID)
	return args.Get(0).(*_playerCoinModel.PlayerCoinShowing), args.Error(1)
}
