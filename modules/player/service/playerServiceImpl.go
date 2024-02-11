package service

import (
	_playerModel "github.com/Rayato159/isekai-shop-api/modules/player/model"
	_playerRepository "github.com/Rayato159/isekai-shop-api/modules/player/repository"
)

type playerServiceImpl struct {
	playerRepository _playerRepository.PlayerRepository
}

func NewPlayerServiceImpl(playerRepository _playerRepository.PlayerRepository) PlayerService {
	return &playerServiceImpl{playerRepository}
}

func (s *playerServiceImpl) CreatePlayer() (string, error) {
	return "", nil
}

func (s *playerServiceImpl) GetPlayerProfile() (*_playerModel.PlayerProfile, error) {
	return nil, nil
}

func (s *playerServiceImpl) EditProfile() error {
	return nil
}
