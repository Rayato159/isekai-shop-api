package repository

import entities "github.com/Rayato159/isekai-shop-api/entities"

type BalancingRepository interface {
	PlayerBalanceRecording(balancingEntity *entities.Balancing) (*entities.Balancing, error)
	PlayerBalanceShowing(playerID string) (*entities.PlayerBalanceDto, error)
}
