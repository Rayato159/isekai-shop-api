package service

import (
	_oauth2Repository "github.com/Rayato159/isekai-shop-api/modules/oauth2/repository"
)

type googleOauth2Service struct {
	oauth2Repository _oauth2Repository.Oauth2Repository
}

func NewGoogleOauth2Service(oauth2Repository _oauth2Repository.Oauth2Repository) Oauth2Service {
	return &googleOauth2Service{oauth2Repository}
}

func (s *googleOauth2Service) UpdateCredential() error {
	return nil
}
