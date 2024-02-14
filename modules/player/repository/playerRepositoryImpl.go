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

func (r *playerRepositoryImpl) FindPlayerById(playerId string) (*_playerEntity.Player, error) {
	player := new(_playerEntity.Player)
	tx := r.db.Where("id = ?", playerId).First(player)

	if tx.Error != nil {
		r.logger.Errorf("Error finding player: %s", tx.Error.Error())
		return nil, &_playerException.FindPlayerException{PlayerID: playerId}
	}

	return player, nil
}

func (r *playerRepositoryImpl) UpdatePlayer(playerID string, updatePlayerDto *_playerEntity.UpdatePlayerDto) (string, error) {
	player := new(_playerEntity.Player)

	tx := r.db.Model(
		player,
	).Where(
		"id = ?", playerID,
	).Updates(
		updatePlayerDto,
	)

	if tx.RowsAffected == 0 {
		r.logger.Errorf("No updating player: %s", tx.Error.Error())
		return "", &_playerException.FindPlayerException{PlayerID: playerID}
	}

	if tx.Error != nil {
		r.logger.Errorf("Error updating player: %s", tx.Error.Error())
		return "", &_playerException.UpdatePlayerException{PlayerID: playerID}
	}

	return playerID, nil
}
