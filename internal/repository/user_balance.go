package repository

import (
	"ewallet/internal/repository/model"
	"ewallet/pkg/connection"
)

type UserBalanceRepository interface {
	GetByUserID(userID string) (*model.UserBalance, error)
	UpdateBalance(balance *model.UserBalance) error
}

type userBalanceRepositorySql struct {
}

func NewUserBalanceRepository() UserBalanceRepository {
	return &userBalanceRepositorySql{}
}

func (r *userBalanceRepositorySql) GetByUserID(userID string) (*model.UserBalance, error) {
	conn := connection.New()
	var balance model.UserBalance
	if err := conn.DB().Where("user_id = ?", userID).First(&balance).Error; err != nil {
		return nil, err
	}
	return &balance, nil
}

func (r *userBalanceRepositorySql) UpdateBalance(balance *model.UserBalance) error {
	conn := connection.New()
	return conn.DB().Save(balance).Error
}
