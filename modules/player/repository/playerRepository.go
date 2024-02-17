package repository

import _playerEntity "github.com/Rayato159/isekai-shop-api/modules/player/entity"

type PlayerRepository interface {
	InsertPlayer(playerEntity *_playerEntity.Player) (*_playerEntity.Player, error)
	FindPlayerByID(playerID string) (*_playerEntity.Player, error)
	UpdatePlayer(playerID string, updatePlayerDto *_playerEntity.UpdatePlayerDto) (string, error)
}
