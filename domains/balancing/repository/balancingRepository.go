package repository

import entities "github.com/Rayato159/isekai-shop-api/domains/entities"

type BalancingRepository interface {
	BalancingRecording(balancingEntity *entities.Balancing) (*entities.Balancing, error)
	PlayerBalanceShowing(playerID string) (*entities.PlayerBalanceDto, error)
}
