package repository

import _playerEntity "github.com/Rayato159/isekai-shop-api/modules/player/entity"

type PlayerRepository interface {
	PlayerCreating(playerEntity *_playerEntity.Player) (*_playerEntity.Player, error)
	FindPlayerByID(playerID string) (*_playerEntity.Player, error)
}
