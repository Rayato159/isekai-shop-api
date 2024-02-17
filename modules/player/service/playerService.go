package service

import (
	_playerModel "github.com/Rayato159/isekai-shop-api/modules/player/model"
)

type PlayerService interface {
	GetPlayer(playerID string) (*_playerModel.Player, error)
	EditPlayer(playerID string, editPlayerReq *_playerModel.EditPlayerReq) (*_playerModel.Player, error)
}
