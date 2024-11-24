package usecase

import (
	"ewallet/internal/repository"
	"ewallet/internal/repository/model"
	"ewallet/internal/usecase/entity"
	"time"

	"github.com/google/uuid"
)

type TopUpUseCase interface {
	TopUp(userID string, amount int) (*entity.TopupResult, error)
}

type topUpUseCaseImpl struct {
	userBalanceRepo repository.UserBalanceRepository
	topupRepo       repository.TopupRepository
}

func NewTopUpUseCase(
	userBalanceRepo repository.UserBalanceRepository,
	topupRepo repository.TopupRepository,
) TopUpUseCase {
	return &topUpUseCaseImpl{
		userBalanceRepo: userBalanceRepo,
		topupRepo:       topupRepo,
	}
}

func (s *topUpUseCaseImpl) TopUp(userID string, amount int) (*entity.TopupResult, error) {
	// Get the user balance
	balance, err := s.userBalanceRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Calculate balance changes
	balanceBefore := balance.Balance
	balanceAfter := balance.Balance + amount

	// Update balance
	balance.Balance = balanceAfter
	if err := s.userBalanceRepo.UpdateBalance(balance); err != nil {
		return nil, err
	}

	// Log the transaction
	transaction := &model.TopUpTransaction{
		TopUpID:       uuid.New().String(),
		BalanceID:     balance.ID,
		Amount:        amount,
		BalanceBefore: balanceBefore,
		BalanceAfter:  balanceAfter,
		CreatedAt:     time.Now(),
	}
	if err := s.topupRepo.Create(transaction); err != nil {
		return nil, err
	}

	return &entity.TopupResult{
		TopUpID:       transaction.TopUpID,
		AmountTopUp:   transaction.Amount,
		BalanceBefore: transaction.BalanceBefore,
		BalanceAfter:  transaction.BalanceAfter,
		CreatedDate:   transaction.CreatedAt,
	}, nil
}
