package service

import (
	_oauth2Model "github.com/Rayato159/isekai-shop-api/modules/oauth2/model"
	_oauth2Repository "github.com/Rayato159/isekai-shop-api/modules/oauth2/repository"
	_playerRepository "github.com/Rayato159/isekai-shop-api/modules/player/repository"
)

type googleOAuth2Service struct {
	oauth2Repository _oauth2Repository.OAuth2Repository
	playerRepository _playerRepository.PlayerRepository
}

func NewGoogleOAuth2Service(
	oauth2Repository _oauth2Repository.OAuth2Repository,
	playerRepository _playerRepository.PlayerRepository,
) OAuth2Service {
	return &googleOAuth2Service{
		oauth2Repository,
		playerRepository,
	}
}

func (s *googleOAuth2Service) ManageUserAccount(createUserInfo *_oauth2Model.CreateUserInfo) error {
	return nil
}
