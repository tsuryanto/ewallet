package repository

import (
	"ewallet/internal/repository/model"
	"ewallet/pkg/connection"
)

type TopupRepository interface {
	Create(transaction *model.TopUpTransaction) error
}

type topupRepositorySql struct {
}

func NewTopupRepository() TopupRepository {
	return &topupRepositorySql{}
}

// UpdateBalanceAndCreateTransaction updates the user's balance and creates a transaction record
func (r *topupRepositorySql) Create(transaction *model.TopUpTransaction) error {
	conn := connection.New()
	return conn.DB().Create(transaction).Error
}
