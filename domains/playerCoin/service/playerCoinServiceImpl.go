package service

import (
	_playerCoinModel "github.com/Rayato159/isekai-shop-api/domains/playerCoin/model"
	_playerCoinRepository "github.com/Rayato159/isekai-shop-api/domains/playerCoin/repository"
	entities "github.com/Rayato159/isekai-shop-api/entities"
)

type playerCoinImpl struct {
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository
}

func NewPlayerCoinServiceImpl(
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository,
) PlayerCoinService {
	return &playerCoinImpl{playerCoinRepository}
}

func (s *playerCoinImpl) BuyingCoin(buyingCoinReq *_playerCoinModel.BuyingCoinReq) (*_playerCoinModel.PlayerCoin, error) {
	playerCoinEntity := &entities.PlayerCoin{
		PlayerID: buyingCoinReq.PlayerID,
		Amount:   buyingCoinReq.Amount,
	}

	insertedPlayerCoin, err := s.playerCoinRepository.Recording(playerCoinEntity)
	if err != nil {
		return nil, err
	}

	return insertedPlayerCoin.ToPlayerCoinModel(), nil
}

func (s *playerCoinImpl) PlayerBalanceShowing(playerID string) *_playerCoinModel.PlayerBalanceShowing {
	balanceDto, err := s.playerCoinRepository.Showing(playerID)
	if err != nil {
		return &_playerCoinModel.PlayerBalanceShowing{
			PlayerID: playerID,
			Balance:  0,
		}
	}

	return balanceDto.ToPlayerBalanceModel()
}