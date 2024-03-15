package repository

import (
	_historyOfPurchasingEntity "github.com/Rayato159/isekai-shop-api/domains/historyOfPurchasing/entity"
)

type HistoryOfPurchasingRepository interface {
	HistoryOfPurchasingRecording(historyOfPurchasingEntity *_historyOfPurchasingEntity.HistoryOfPurchasing) (*_historyOfPurchasingEntity.HistoryOfPurchasing, error)
	PlayerHistoryOfPurchasingListing(playerID string) ([]*_historyOfPurchasingEntity.HistoryOfPurchasing, error)
}
