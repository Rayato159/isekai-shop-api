package service

import (
	_playerModel "github.com/Rayato159/isekai-shop-api/modules/player/model"
)

type PlayerService interface {
	GetPlayerProfile(playerID string) (*_playerModel.PlayerProfile, error)
	EditPlayerProfile(playerID string, updatePlayer *_playerModel.UpdatePlayerProfile) (*_playerModel.PlayerProfile, error)
}
