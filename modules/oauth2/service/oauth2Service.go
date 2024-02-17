package service

import (
	_adminModel "github.com/Rayato159/isekai-shop-api/modules/admin/model"
	_playerModel "github.com/Rayato159/isekai-shop-api/modules/player/model"
)

type OAuth2Service interface {
	CreatePlayerAccount(createPlayerReq *_playerModel.CreatePlayerReq) error
	CreateAdminAccount(createAdminReq *_adminModel.CreateAdminReq) error
}
