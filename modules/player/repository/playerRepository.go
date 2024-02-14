package repository

import _playerEntity "github.com/Rayato159/isekai-shop-api/modules/player/entity"

type PlayerRepository interface {
	InsertPlayer(playerEntitiy *_playerEntity.Player) (string, error)
	FindPlayerById(playerId string) (*_playerEntity.Player, error)
}