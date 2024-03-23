package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"
	_player "github.com/Rayato159/isekai-shop-api/pkg/player/exception"

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

func (r *playerRepositoryImpl) Creating(playerEntity *entities.Player) (*entities.Player, error) {
	insertedPlayer := new(entities.Player)

	if err := r.db.Create(playerEntity).Scan(insertedPlayer).Error; err != nil {
		r.logger.Error("Creating player failed", err.Error())
		return nil, &_player.PlayerCreating{PlayerID: playerEntity.ID}
	}

	return insertedPlayer, nil
}

func (r *playerRepositoryImpl) FindByID(playerID string) (*entities.Player, error) {
	player := new(entities.Player)
	tx := r.db.Where("id = ?", playerID).First(player)

	if tx.Error != nil {
		r.logger.Errorf("Finding player failed: %s", tx.Error.Error())
		return nil, &_player.PlayerNotFound{PlayerID: playerID}
	}

	return player, nil
}
