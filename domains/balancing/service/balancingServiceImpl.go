package service

import (
	_balancingModel "github.com/Rayato159/isekai-shop-api/domains/balancing/model"
	_balancingRepository "github.com/Rayato159/isekai-shop-api/domains/balancing/repository"
	entities "github.com/Rayato159/isekai-shop-api/entities"
)

type balancingServiceImpl struct {
	balancingRepository _balancingRepository.BalancingRepository
}

func NewBalancingServiceImpl(
	balancingRepository _balancingRepository.BalancingRepository,
) BalancingService {
	return &balancingServiceImpl{balancingRepository}
}

func (s *balancingServiceImpl) TopUp(topUpReq *_balancingModel.TopUpReq) (*_balancingModel.Balancing, error) {
	balancingEntity := &entities.Balancing{
		PlayerID: topUpReq.PlayerID,
		Amount:   topUpReq.Amount,
	}

	insertedBalancing, err := s.balancingRepository.BalancingRecording(balancingEntity)
	if err != nil {
		return nil, err
	}

	return insertedBalancing.ToBalancingModel(), nil
}

func (s *balancingServiceImpl) PlayerBalanceShowing(playerID string) *_balancingModel.PlayerBalance {
	balanceDto, err := s.balancingRepository.PlayerBalanceShowing(playerID)
	if err != nil {
		return &_balancingModel.PlayerBalance{
			PlayerID: playerID,
			Balance:  0,
		}
	}

	return balanceDto.ToPlayerBalanceModel()
}
