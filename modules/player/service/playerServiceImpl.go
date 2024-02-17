package service

import (
	_playerEntity "github.com/Rayato159/isekai-shop-api/modules/player/entity"
	_playerModel "github.com/Rayato159/isekai-shop-api/modules/player/model"
	_playerRepository "github.com/Rayato159/isekai-shop-api/modules/player/repository"
	"github.com/labstack/echo/v4"
)

type playerServiceImpl struct {
	playerRepository _playerRepository.PlayerRepository
	logger           echo.Logger
}

func NewPlayerServiceImpl(playerRepository _playerRepository.PlayerRepository, logger echo.Logger) PlayerService {
	return &playerServiceImpl{
		playerRepository: playerRepository,
		logger:           logger,
	}
}

func (s *playerServiceImpl) GetPlayer(playerID string) (*_playerModel.Player, error) {
	player, err := s.playerRepository.FindPlayerByID(playerID)
	if err != nil {
		return nil, err
	}

	return &_playerModel.Player{
		ID:       player.ID,
		Username: player.Username,
		Email:    player.Email,
		Name:     player.Name,
		Avatar:   player.Avatar,
	}, nil
}

func (s *playerServiceImpl) EditPlayer(playerID string, editPlayerReq *_playerModel.EditPlayerReq) (*_playerModel.Player, error) {
	editPlayerReqDto := &_playerEntity.UpdatePlayerDto{
		Username: editPlayerReq.Username,
	}

	updatedPlayerID, err := s.playerRepository.UpdatePlayer(playerID, editPlayerReqDto)
	if err != nil {
		return nil, err
	}

	playerEntitiy, err := s.playerRepository.FindPlayerByID(updatedPlayerID)
	if err != nil {
		return nil, err
	}

	return playerEntitiy.ToPlayerModel(), nil
}
