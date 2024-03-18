package repository

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"
)

type PurchasingRepository interface {
	PurchasingHistoryRecording(purchasingEntity *entities.PurchasingHistory) (*entities.PurchasingHistory, error)
}
