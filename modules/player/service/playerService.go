package service

import (
	_playerModel "github.com/Rayato159/isekai-shop-api/modules/player/model"
)

type PlayerService interface {
	PlayerInventoryListing(playerID string) ([]*_playerModel.Inventory, error)
}
