package repository

import entities "github.com/Rayato159/isekai-shop-api/domains/entities"

type PlayerRepository interface {
	PlayerCreating(playerEntity *entities.Player) (*entities.Player, error)
	FindPlayerByID(playerID string) (*entities.Player, error)
}
