package service

import (
	_oauth2Model "github.com/Rayato159/isekai-shop-api/modules/oauth2/model"
)

type OAuth2Service interface {
	CreatePlayerAccount(createPlayerInfo *_oauth2Model.CreatePlayerInfo) error
}
