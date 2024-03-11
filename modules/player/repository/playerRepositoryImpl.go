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

func (r *playerRepositoryImpl) PlayerCreating(playerEntity *_playerEntity.Player) (*_playerEntity.Player, error) {
	insertedPlayer := new(_playerEntity.Player)

	if err := r.db.Create(playerEntity).Scan(insertedPlayer).Error; err != nil {
		r.logger.Error("Failed to insert item", err.Error())
		return nil, &_playerException.PlayerCreatingException{PlayerID: playerEntity.ID}
	}

	return insertedPlayer, nil
}

func (r *playerRepositoryImpl) FindPlayerByID(playerID string) (*_playerEntity.Player, error) {
	player := new(_playerEntity.Player)
	tx := r.db.Where("id = ?", playerID).First(player)

	if tx.Error != nil {
		r.logger.Errorf("Error finding player: %s", tx.Error.Error())
		return nil, &_playerException.PlayerNotFoundException{PlayerID: playerID}
	}

	return player, nil
}

func (r *playerRepositoryImpl) ProfileEditing(playerID string, updatePlayerDto *_playerEntity.ProfileEditingDto) (string, error) {
	updatedPlayer := new(_playerEntity.Player)

	tx := r.db.Model(&_playerEntity.Player{}).Where(
		"id = ?", playerID,
	).Updates(
		updatePlayerDto,
	).Scan(updatedPlayer)

	if tx.Error != nil {
		r.logger.Errorf("Error updating player: %s", tx.Error.Error())
		return "", &_playerException.ProfileEditingException{PlayerID: playerID}
	}

	return playerID, nil
}
