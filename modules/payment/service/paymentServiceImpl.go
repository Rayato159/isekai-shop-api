package service

import (
	_paymentEntity "github.com/Rayato159/isekai-shop-api/modules/payment/entity"
	_paymentModel "github.com/Rayato159/isekai-shop-api/modules/payment/model"
	_paymentRepository "github.com/Rayato159/isekai-shop-api/modules/payment/repository"
	"github.com/labstack/echo/v4"
)

type paymentServiceImpl struct {
	paymentRepository _paymentRepository.PaymentRepository
	logger            echo.Logger
}

func NewPaymentServiceImpl(paymentRepository _paymentRepository.PaymentRepository, logger echo.Logger) PaymentService {
	return &paymentServiceImpl{
		paymentRepository: paymentRepository,
		logger:            logger,
	}
}

func (s *paymentServiceImpl) TopUp(topUpReq *_paymentModel.TopUpReq) (*_paymentModel.Payment, error) {
	paymentEntity := &_paymentEntity.Payment{
		PlayerID: topUpReq.PlayerID,
		Amount:   topUpReq.Amount,
	}

	insertedPayment, err := s.paymentRepository.TopUp(paymentEntity)
	if err != nil {
		return nil, err
	}

	return insertedPayment.ToPaymentModel(), nil
}

func (s *paymentServiceImpl) CalculatePlayerBalance(playerID string) *_paymentModel.PlayerBalance {
	balanceDto, err := s.paymentRepository.CalculatePlayerBalance(playerID)
	if err != nil {
		return &_paymentModel.PlayerBalance{
			PlayerID: playerID,
			Balance:  0,
		}
	}

	return balanceDto.ToPlayerBalanceModel()
}
