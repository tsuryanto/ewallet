package reqres

import "time"

// TopUpRequest represents the request body for a top-up
type TopUpRequest struct {
	Amount int `json:"amount"`
}

// TopUpResponse represents the response body for a top-up
type TopUpResponse struct {
	TopUpID       string    `json:"top_up_id"`
	AmountTopUp   int       `json:"amount_top_up"`
	BalanceBefore int       `json:"balance_before"`
	BalanceAfter  int       `json:"balance_after"`
	CreatedDate   time.Time `json:"created_date"`
}
