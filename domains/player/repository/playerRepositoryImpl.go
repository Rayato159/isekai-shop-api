package repository

import (
	_playerException "github.com/Rayato159/isekai-shop-api/domains/player/exception"
	entities "github.com/Rayato159/isekai-shop-api/entities"

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

func (r *playerRepositoryImpl) PlayerCreating(playerEntity *entities.Player) (*entities.Player, error) {
	insertedPlayer := new(entities.Player)

	if err := r.db.Create(playerEntity).Scan(insertedPlayer).Error; err != nil {
		r.logger.Error("Failed to insert item", err.Error())
		return nil, &_playerException.PlayerCreatingException{PlayerID: playerEntity.ID}
	}

	return insertedPlayer, nil
}

func (r *playerRepositoryImpl) FindPlayerByID(playerID string) (*entities.Player, error) {
	player := new(entities.Player)
	tx := r.db.Where("id = ?", playerID).First(player)

	if tx.Error != nil {
		r.logger.Errorf("Error finding player: %s", tx.Error.Error())
		return nil, &_playerException.PlayerNotFoundException{PlayerID: playerID}
	}

	return player, nil
}
