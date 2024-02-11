package repository

import (
	_playerEntity "github.com/Rayato159/isekai-shop-api/modules/player/entity"
	"gorm.io/gorm"
)

type playerRepositoryImpl struct {
	db *gorm.DB
}

func NewPlayerRepositoryImpl(db *gorm.DB) PlayerRepository {
	return &playerRepositoryImpl{db}
}

func (r *playerRepositoryImpl) InsertPlayer(playerEntitiy *_playerEntity.Player) (string, error) {
	return "", nil
}
