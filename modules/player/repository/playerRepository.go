package repository

import _playerEntity "github.com/Rayato159/isekai-shop-api/modules/player/entity"

type PlayerRepository interface {
	InsertPlayer(playerEntitiy *_playerEntity.Player) (string, error)
	FindPlayerByID(playerID string) (*_playerEntity.Player, error)
	UpdatePlayer(playerID string, updatePlayerDto *_playerEntity.UpdatePlayerDto) (*_playerEntity.Player, error)
}
