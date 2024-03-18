package repository

import (
	_purchasingEntity "github.com/Rayato159/isekai-shop-api/domains/purchasing/entity"
)

type PurchasingRepository interface {
	PurchasingHistoryRecording(purchasingEntity *_purchasingEntity.Purchasing) (*_purchasingEntity.Purchasing, error)
}
