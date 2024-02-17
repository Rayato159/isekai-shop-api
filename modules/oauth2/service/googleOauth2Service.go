package service

import (
	_adminEntity "github.com/Rayato159/isekai-shop-api/modules/admin/entity"
	_adminRepository "github.com/Rayato159/isekai-shop-api/modules/admin/repository"
	_oauth2Model "github.com/Rayato159/isekai-shop-api/modules/oauth2/model"
	_playerEntity "github.com/Rayato159/isekai-shop-api/modules/player/entity"
	_playerRepository "github.com/Rayato159/isekai-shop-api/modules/player/repository"
	"github.com/labstack/echo/v4"
)

type googleOAuth2Service struct {
	playerRepository _playerRepository.PlayerRepository
	adminRepository  _adminRepository.AdminRepository
	logger           echo.Logger
}

func NewGoogleOAuth2Service(
	playerRepository _playerRepository.PlayerRepository,
	adminRepository _adminRepository.AdminRepository,
	logger echo.Logger,
) OAuth2Service {
	return &googleOAuth2Service{
		playerRepository: playerRepository,
		adminRepository:  adminRepository,
		logger:           logger,
	}
}

func (s *googleOAuth2Service) CreatePlayerAccount(createPlayerInfo *_oauth2Model.CreatePlayerInfo) error {
	if !s.isPlayerIsExists(createPlayerInfo.ID) {
		playerEntity := &_playerEntity.Player{
			ID:     createPlayerInfo.ID,
			Email:  createPlayerInfo.Email,
			Name:   createPlayerInfo.Name,
			Avatar: createPlayerInfo.Avatar,
		}

		_, err := s.playerRepository.InsertPlayer(playerEntity)
		if err != nil {
			return err
		}
	}

	s.logger.Infof("Player created: %s", createPlayerInfo.ID)

	return nil
}

func (s *googleOAuth2Service) CreateAdminAccount(createAdminInfo *_oauth2Model.CreateAdminInfo) error {
	if !s.isAdminIsExists(createAdminInfo.ID) {
		adminEntity := &_adminEntity.Admin{
			ID:     createAdminInfo.ID,
			Email:  createAdminInfo.Email,
			Name:   createAdminInfo.Name,
			Avatar: createAdminInfo.Avatar,
		}

		_, err := s.adminRepository.InsertAdmin(adminEntity)
		if err != nil {
			return err
		}
	}

	s.logger.Infof("Admin created: %s", createAdminInfo.ID)

	return nil

}

func (s *googleOAuth2Service) isPlayerIsExists(palyerId string) bool {
	player, err := s.playerRepository.FindPlayerByID(palyerId)
	if err != nil {
		return false
	}

	return player != nil
}

func (s *googleOAuth2Service) isAdminIsExists(adminId string) bool {
	admin, err := s.adminRepository.FindAdminByID(adminId)
	if err != nil {
		return false
	}

	return admin != nil
}
