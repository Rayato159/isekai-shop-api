package repository

import "gorm.io/gorm"

type googleOauth2Repository struct {
	db *gorm.DB
}

func NewGoogleOauth2Repository(db *gorm.DB) Oauth2Repository {
	return &googleOauth2Repository{db}
}

func (r *googleOauth2Repository) InsertOauth2() error {
	return nil
}
