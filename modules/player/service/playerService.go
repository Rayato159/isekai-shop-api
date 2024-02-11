package service

import (
	_playerModel "github.com/Rayato159/isekai-shop-api/modules/player/model"
)

type PlayerService interface {
	CreatePlayer() (string, error)
	EditProfile() error
	GetPlayerProfile() (*_playerModel.PlayerProfile, error)
}
