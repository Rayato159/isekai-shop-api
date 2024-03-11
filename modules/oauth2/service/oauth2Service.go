package service

import (
	_adminModel "github.com/Rayato159/isekai-shop-api/modules/admin/model"
	_playerModel "github.com/Rayato159/isekai-shop-api/modules/player/model"
)

type OAuth2Service interface {
	CreatePlayerAccount(playerCreatingReq *_playerModel.CreatePlayerReq) error
	CreateAdminAccount(createAdminReq *_adminModel.CreateAdminReq) error
	IsThisGuyIsReallyPlayer(playerID string) bool
	IsThisGuyIsReallyAdmin(adminID string) bool
}
