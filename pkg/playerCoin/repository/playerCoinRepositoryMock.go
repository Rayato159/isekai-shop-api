package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"

	"github.com/stretchr/testify/mock"
)

type CoinRepositoryMock struct {
	mock.Mock
}

func (m *CoinRepositoryMock) Recording(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error) {
	args := m.Called(playerCoinEntity)
	return args.Get(0).(*entities.PlayerCoin), args.Error(1)
}

func (m *CoinRepositoryMock) Showing(playerID string) (*entities.PlayerCoinShowingDto, error) {
	args := m.Called(playerID)
	return args.Get(0).(*entities.PlayerCoinShowingDto), args.Error(1)
}
