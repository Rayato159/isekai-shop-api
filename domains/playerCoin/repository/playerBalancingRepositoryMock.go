package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"

	"github.com/stretchr/testify/mock"
)

type BalancingRepositoryMock struct {
	mock.Mock
}

func (m *BalancingRepositoryMock) Recording(balancingEntity *entities.PlayerCoin) (*entities.PlayerCoin, error) {
	args := m.Called(balancingEntity)
	return args.Get(0).(*entities.PlayerCoin), args.Error(1)
}

func (m *BalancingRepositoryMock) Showing(playerID string) (*entities.PlayerBalanceShowingDto, error) {
	args := m.Called(playerID)
	return args.Get(0).(*entities.PlayerBalanceShowingDto), args.Error(1)
}
