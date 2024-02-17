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

func (r *playerRepositoryImpl) FindPlayerByID(playerID string) (*_playerEntity.Player, error) {
	player := new(_playerEntity.Player)
	tx := r.db.Where("id = ?", playerID).First(player)

	if tx.Error != nil {
		r.logger.Errorf("Error finding player: %s", tx.Error.Error())
		return nil, &_playerException.PlayerNotFoundException{PlayerID: playerID}
	}

	return player, nil
}

func (r *playerRepositoryImpl) UpdatePlayer(playerID string, updatePlayerDto *_playerEntity.UpdatePlayerDto) (string, error) {
	updatedPlayer := new(_playerEntity.Player)

	tx := r.db.Model(&_playerEntity.Player{}).Where(
		"id = ?", playerID,
	).Updates(
		updatePlayerDto,
	).Scan(updatedPlayer)

	if tx.Error != nil {
		r.logger.Errorf("Error updating player: %s", tx.Error.Error())
		return "", &_playerException.UpdatePlayerException{PlayerID: playerID}
	}

	return playerID, nil
}
