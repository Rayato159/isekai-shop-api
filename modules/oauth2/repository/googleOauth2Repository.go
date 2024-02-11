package repository

import (
	"gorm.io/gorm"
)

type googleOAuth2Repository struct {
	db *gorm.DB
}

func NewGoogleOAuth2Repository(db *gorm.DB) OAuth2Repository {
	return &googleOAuth2Repository{db}
}

func (r *googleOAuth2Repository) InsertOAuth2() error {
	return nil
}
