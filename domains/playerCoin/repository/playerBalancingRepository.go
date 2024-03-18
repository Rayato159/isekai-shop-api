package repository

import entities "github.com/Rayato159/isekai-shop-api/entities"

type PlayerCoinRepository interface {
	Recording(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
	Showing(playerID string) (*entities.PlayerBalanceShowingDto, error)
}
