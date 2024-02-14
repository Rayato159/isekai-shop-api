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

func (s *playerServiceImpl) GetPlayerProfile(playerID string) (*_playerModel.PlayerProfile, error) {
	player, err := s.playerRepository.FindPlayerById(playerID)
	if err != nil {
		return nil, err
	}

	return &_playerModel.PlayerProfile{
		ID:       player.ID,
		Username: player.Username,
		Email:    player.Email,
		Name:     player.Name,
		Avatar:   player.Avatar,
	}, nil
}

func (s *playerServiceImpl) EditPlayerProfile(playerID string, updatePlayer *_playerModel.UpdatePlayerProfile) (*_playerModel.PlayerProfile, error) {
	updatePlayerEntity := &_playerEntity.UpdatePlayer{
		Username: updatePlayer.Username,
	}

	_, err := s.playerRepository.UpdatePlayer(playerID, updatePlayerEntity)
	if err != nil {
		return nil, err
	}

	player, err := s.playerRepository.FindPlayerById(playerID)
	if err != nil {
		return nil, err
	}

	return &_playerModel.PlayerProfile{
		ID:       player.ID,
		Username: player.Username,
		Email:    player.Email,
		Name:     player.Name,
		Avatar:   player.Avatar,
	}, nil
}
