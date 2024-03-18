package repository

import entities "github.com/Rayato159/isekai-shop-api/entities"

type PlayerBalancingRepository interface {
	Recording(playerBalancingEntity *entities.PlayerBalancing) (*entities.PlayerBalancing, error)
	Showing(playerID string) (*entities.PlayerBalanceShowingDto, error)
}
