package repository

import entities "github.com/Rayato159/isekai-shop-api/entities"

type PlayerRepository interface {
	Creating(playerEntity *entities.Player) (*entities.Player, error)
	FindByID(playerID string) (*entities.Player, error)
}
