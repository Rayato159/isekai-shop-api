package service

import (
	_oauth2Model "github.com/Rayato159/isekai-shop-api/modules/oauth2/model"
	_playerEntity "github.com/Rayato159/isekai-shop-api/modules/player/entity"
	_playerRepository "github.com/Rayato159/isekai-shop-api/modules/player/repository"
	"github.com/labstack/echo/v4"
)

type googleOAuth2Service struct {
	playerRepository _playerRepository.PlayerRepository
	logger           echo.Logger
}

func NewGoogleOAuth2Service(
	playerRepository _playerRepository.PlayerRepository,
	logger echo.Logger,
) OAuth2Service {
	return &googleOAuth2Service{
		playerRepository: playerRepository,
		logger:           logger,
	}
}

func (s *googleOAuth2Service) CreatePlayerAccount(createPlayerInfo *_oauth2Model.CreatePlayerInfo) error {
	if !s.isPlayerIsExists(createPlayerInfo.ID) {
		playerEntity := &_playerEntity.Player{
			ID:     createPlayerInfo.ID,
			Email:  createPlayerInfo.Email,
			Name:   createPlayerInfo.Name,
			Avatar: createPlayerInfo.Picture,
		}

		_, err := s.playerRepository.InsertPlayer(playerEntity)
		if err != nil {
			return err
		}
	}

	s.logger.Infof("Player created: %s", createPlayerInfo.ID)

	return nil
}

func (s *googleOAuth2Service) isPlayerIsExists(palyerId string) bool {
	player, err := s.playerRepository.FindPlayerById(palyerId)
	if err != nil {
		return false
	}

	return player != nil
}
