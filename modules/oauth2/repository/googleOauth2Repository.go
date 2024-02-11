package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_oauth2Entity "github.com/Rayato159/isekai-shop-api/modules/oauth2/entity"
	_oauth2Exception "github.com/Rayato159/isekai-shop-api/modules/oauth2/exception"
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

func (r *googleOAuth2Repository) InsertPassport(passportEntity *_oauth2Entity.Passport) error {
	tx := r.db.Create(passportEntity)

	if tx.Error != nil {
		r.logger.Errorf("Error inserting passport: %s", tx.Error.Error())
		return &_oauth2Exception.InsertPassportException{PlayerID: passportEntity.PlayerID}
	}

	return nil
}
