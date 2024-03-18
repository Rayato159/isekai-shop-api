package repository

import _balancingEntity "github.com/Rayato159/isekai-shop-api/domains/balancing/entity"

type BalancingRepository interface {
	BalancingRecording(balancingEntity *_balancingEntity.Balancing) (*_balancingEntity.Balancing, error)
	PlayerBalanceShowing(playerID string) (*_balancingEntity.PlayerBalanceDto, error)
}
