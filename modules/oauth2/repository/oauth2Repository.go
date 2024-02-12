package repository

import (
	_oauth2Entity "github.com/Rayato159/isekai-shop-api/modules/oauth2/entity"
)

type OAuth2Repository interface {
	InsertPassport(passportEntity *_oauth2Entity.Passport) error
	DeletePassport(refreshToken string) error
}
