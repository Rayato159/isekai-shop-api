package service

import (
	_oauth2Model "github.com/Rayato159/isekai-shop-api/modules/oauth2/model"
)

type OAuth2Service interface {
	ManageUserAccount(createUserInfo *_oauth2Model.CreateUserInfo) error
}
