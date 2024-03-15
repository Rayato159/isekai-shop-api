package service

import (
	_playerModel "github.com/Rayato159/isekai-shop-api/domains/player/model"
)

type PlayerService interface {
	PlayerInventoryListing(playerID string) ([]*_playerModel.Inventory, error)
}
