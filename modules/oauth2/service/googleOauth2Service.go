package service

import (
	_adminEntity "github.com/Rayato159/isekai-shop-api/modules/admin/entity"
	_adminModel "github.com/Rayato159/isekai-shop-api/modules/admin/model"
	_adminRepository "github.com/Rayato159/isekai-shop-api/modules/admin/repository"
	_playerEntity "github.com/Rayato159/isekai-shop-api/modules/player/entity"
	_playerModel "github.com/Rayato159/isekai-shop-api/modules/player/model"
	_playerRepository "github.com/Rayato159/isekai-shop-api/modules/player/repository"
)

type googleOAuth2Service struct {
	playerRepository _playerRepository.PlayerRepository
	adminRepository  _adminRepository.AdminRepository
}

func NewGoogleOAuth2Service(
	playerRepository _playerRepository.PlayerRepository,
	adminRepository _adminRepository.AdminRepository,
) OAuth2Service {
	return &googleOAuth2Service{
		playerRepository: playerRepository,
		adminRepository:  adminRepository,
	}
}

func (s *googleOAuth2Service) CreatePlayerAccount(playerCreatingReq *_playerModel.CreatePlayerReq) error {
	if !s.isPlayerIsExists(playerCreatingReq.ID) {
		playerEntity := &_playerEntity.Player{
			ID:     playerCreatingReq.ID,
			Email:  playerCreatingReq.Email,
			Name:   playerCreatingReq.Name,
			Avatar: playerCreatingReq.Avatar,
		}

		_, err := s.playerRepository.InsertPlayer(playerEntity)
		if err != nil {
			return err
		}
	}

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
