package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type googleOAuth2Repository struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewGoogleOAuth2Repository(db *gorm.DB, logger echo.Logger) OAuth2Repository {
	return &googleOAuth2Repository{
		db:     db,
		logger: logger,
	}
}

func (r *googleOAuth2Repository) InsertOAuth2() error {
	return nil
}
