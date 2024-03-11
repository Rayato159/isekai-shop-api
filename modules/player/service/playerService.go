package service

import (
	_playerModel "github.com/Rayato159/isekai-shop-api/modules/player/model"
)

type PlayerService interface {
	PlayerProfiling(playerID string) (*_playerModel.Player, error)
	PlayerProfileEditing(playerID string, editPlayerReq *_playerModel.PlayerProfileEditingReq) (*_playerModel.Player, error)
}
