package service

import (
	entities "github.com/Rayato159/isekai-shop-api/entities"
	_playerCoinModel "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/model"
	_playerCoinRepository "github.com/Rayato159/isekai-shop-api/pkg/playerCoin/repository"
)

type playerCoinImpl struct {
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository
}

func NewPlayerCoinServiceImpl(
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository,
) PlayerCoinService {
	return &playerCoinImpl{playerCoinRepository}
}

func (s *playerCoinImpl) CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin, error) {
	playerCoinEntity := &entities.PlayerCoin{
		PlayerID: coinAddingReq.PlayerID,
		Amount:   coinAddingReq.Amount,
	}

	insertedPlayerCoin, err := s.playerCoinRepository.Recording(playerCoinEntity)
	if err != nil {
		return nil, err
	}

	return insertedPlayerCoin.ToPlayerCoinModel(), nil
}

func (s *playerCoinImpl) PlayerCoinShowing(playerID string) *_playerCoinModel.PlayerCoinShowing {
	coin, err := s.playerCoinRepository.Showing(playerID)
	if err != nil {
		return &_playerCoinModel.PlayerCoinShowing{
			PlayerID: playerID,
			Coin:     0,
		}
	}

	return coin
}
