package entity

import "time"

type TopupResult struct {
	TopUpID       string
	AmountTopUp   int
	BalanceBefore int
	BalanceAfter  int
	CreatedDate   time.Time
}

func NewTopupResult(topUpID string, amountTopUp int, balanceBefore int, balanceAfter int, createdDate time.Time) TopupResult {
	return TopupResult{
		TopUpID:       topUpID,
		AmountTopUp:   amountTopUp,
		BalanceBefore: balanceBefore,
		BalanceAfter:  balanceAfter,
		CreatedDate:   createdDate,
	}
}
