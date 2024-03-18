package repository

import (
	_purchasingEntity "github.com/Rayato159/isekai-shop-api/domains/purchasing/entity"
)

type PurchasingRepository interface {
	PurchasingHistoryRecording(purchasingEntity *_purchasingEntity.PurchasingHistory) (*_purchasingEntity.PurchasingHistory, error)
}
