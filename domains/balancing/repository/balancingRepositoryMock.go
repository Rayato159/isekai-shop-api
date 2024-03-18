package repository

import (
	_balancingEntity "github.com/Rayato159/isekai-shop-api/domains/balancing/entity"

	"github.com/stretchr/testify/mock"
)

type BalancingRepositoryMock struct {
	mock.Mock
}

func (m *BalancingRepositoryMock) BalancingRecording(balancingEntity *_balancingEntity.Balancing) (*_balancingEntity.Balancing, error) {
	args := m.Called(balancingEntity)
	return args.Get(0).(*_balancingEntity.Balancing), args.Error(1)
}

func (m *BalancingRepositoryMock) PlayerBalanceShowing(playerID string) (*_balancingEntity.PlayerBalanceDto, error) {
	args := m.Called(playerID)
	return args.Get(0).(*_balancingEntity.PlayerBalanceDto), args.Error(1)
}
