package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"

	"github.com/stretchr/testify/mock"
)

type BalancingRepositoryMock struct {
	mock.Mock
}

func (m *BalancingRepositoryMock) PlayerBalanceRecording(balancingEntity *entities.Balancing) (*entities.Balancing, error) {
	args := m.Called(balancingEntity)
	return args.Get(0).(*entities.Balancing), args.Error(1)
}

func (m *BalancingRepositoryMock) PlayerBalanceShowing(playerID string) (*entities.PlayerBalanceDto, error) {
	args := m.Called(playerID)
	return args.Get(0).(*entities.PlayerBalanceDto), args.Error(1)
}
