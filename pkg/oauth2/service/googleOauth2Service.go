package service

import (
	_adminModel "github.com/Rayato159/isekai-shop-api/pkg/admin/model"
	_adminRepository "github.com/Rayato159/isekai-shop-api/pkg/admin/repository"
	_playerModel "github.com/Rayato159/isekai-shop-api/pkg/player/model"
	_playerRepository "github.com/Rayato159/isekai-shop-api/pkg/player/repository"
	entities "github.com/Rayato159/isekai-shop-api/entities"
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

func (s *googleOAuth2Service) PlayerAccountCreating(playerCreatingReq *_playerModel.CreatePlayerReq) error {
	if !s.isPlayerIsExists(playerCreatingReq.ID) {
		playerEntity := &entities.Player{
			ID:     playerCreatingReq.ID,
			Email:  playerCreatingReq.Email,
			Name:   playerCreatingReq.Name,
			Avatar: playerCreatingReq.Avatar,
		}

		_, err := s.playerRepository.Creating(playerEntity)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *googleOAuth2Service) AdminAccountCreating(createAdminInfo *_adminModel.CreateAdminReq) error {
	if !s.isAdminIsExists(createAdminInfo.ID) {
		adminEntity := &entities.Admin{
			ID:     createAdminInfo.ID,
			Email:  createAdminInfo.Email,
			Name:   createAdminInfo.Name,
			Avatar: createAdminInfo.Avatar,
		}

		_, err := s.adminRepository.Creating(adminEntity)
		if err != nil {
			return err
		}
	}

	return nil

}

func (s *googleOAuth2Service) IsThisGuyIsReallyAdmin(adminID string) bool {
	if _, err := s.adminRepository.FindByID(adminID); err != nil {
		return false
	}
	return true
}

func (s *googleOAuth2Service) IsThisGuyIsReallyPlayer(playerID string) bool {
	if _, err := s.playerRepository.FindByID(playerID); err != nil {
		return false
	}
	return true
}

func (s *googleOAuth2Service) isPlayerIsExists(palyerId string) bool {
	player, err := s.playerRepository.FindByID(palyerId)
	if err != nil {
		return false
	}

	return player != nil
}

func (s *googleOAuth2Service) isAdminIsExists(adminId string) bool {
	admin, err := s.adminRepository.FindByID(adminId)
	if err != nil {
		return false
	}

	return admin != nil
}
