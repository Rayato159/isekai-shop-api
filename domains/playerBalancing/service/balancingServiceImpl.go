package service

import (
	_playerBalancingModel "github.com/Rayato159/isekai-shop-api/domains/playerBalancing/model"
	_playerBalancingRepository "github.com/Rayato159/isekai-shop-api/domains/playerBalancing/repository"
	entities "github.com/Rayato159/isekai-shop-api/entities"
)

type playerBalancingImpl struct {
	playerBalancingRepository _playerBalancingRepository.PlayerBalancingRepository
}

func NewPlayerBalancingServiceImpl(
	playerBalancingRepository _playerBalancingRepository.PlayerBalancingRepository,
) PlayerBalancingService {
	return &playerBalancingImpl{playerBalancingRepository}
}

func (s *playerBalancingImpl) TopUp(topUpReq *_playerBalancingModel.TopUpReq) (*_playerBalancingModel.PlayerBalancing, error) {
	playerBalancingEntity := &entities.PlayerBalancing{
		PlayerID: topUpReq.PlayerID,
		Amount:   topUpReq.Amount,
	}

	insertedPlayerBalancing, err := s.playerBalancingRepository.Recording(playerBalancingEntity)
	if err != nil {
		return nil, err
	}

	return insertedPlayerBalancing.ToPlayerBalancingModel(), nil
}

func (s *playerBalancingImpl) PlayerBalanceShowing(playerID string) *_playerBalancingModel.PlayerBalanceShowing {
	balanceDto, err := s.playerBalancingRepository.Showing(playerID)
	if err != nil {
		return &_playerBalancingModel.PlayerBalanceShowing{
			PlayerID: playerID,
			Balance:  0,
		}
	}

	return balanceDto.ToPlayerBalanceModel()
}
