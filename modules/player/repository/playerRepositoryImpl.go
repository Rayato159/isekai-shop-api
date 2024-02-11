package repository

import (
	_playerEntity "github.com/Rayato159/isekai-shop-api/modules/player/entity"
	_playerException "github.com/Rayato159/isekai-shop-api/modules/player/exception"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type playerRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewPlayerRepositoryImpl(db *gorm.DB, logger echo.Logger) PlayerRepository {
	return &playerRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *playerRepositoryImpl) InsertPlayer(playerEntitiy *_playerEntity.Player) (string, error) {
	tx := r.db.Create(playerEntitiy)

	if tx.Error != nil {
		r.logger.Errorf("Error inserting player: %s", tx.Error.Error())
		return "", &_playerException.InsertPlayerException{PlayerID: playerEntitiy.ID}
	}

	return playerEntitiy.ID, nil
}
