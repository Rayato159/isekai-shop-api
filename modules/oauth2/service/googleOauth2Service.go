package service

import (
	_adminEntity "github.com/Rayato159/isekai-shop-api/modules/admin/entity"
	_adminModel "github.com/Rayato159/isekai-shop-api/modules/admin/model"
	_adminRepository "github.com/Rayato159/isekai-shop-api/modules/admin/repository"
	_playerEntity "github.com/Rayato159/isekai-shop-api/modules/player/entity"
	_playerModel "github.com/Rayato159/isekai-shop-api/modules/player/model"
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

func (s *googleOAuth2Service) CreatePlayerAccount(createPlayerReq *_playerModel.CreatePlayerReq) error {
	if !s.isPlayerIsExists(createPlayerReq.ID) {
		playerEntity := &_playerEntity.Player{
			ID:     createPlayerReq.ID,
			Email:  createPlayerReq.Email,
			Name:   createPlayerReq.Name,
			Avatar: createPlayerReq.Avatar,
		}

		_, err := s.playerRepository.InsertPlayer(playerEntity)
		if err != nil {
			return err
		}
	}

	s.logger.Infof("Player created: %s", createPlayerReq.ID)

	return nil
}

func (s *googleOAuth2Service) CreateAdminAccount(createAdminInfo *_adminModel.CreateAdminReq) error {
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

func (s *googleOAuth2Service) IsThisGuyIsReallyAdmin(adminID string) bool {
	if _, err := s.adminRepository.FindAdminByID(adminID); err != nil {
		return false
	}
	return true
}

func (s *googleOAuth2Service) IsThisGuyIsReallyPlayer(playerID string) bool {
	if _, err := s.playerRepository.FindPlayerByID(playerID); err != nil {
		return false
	}
	return true
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
