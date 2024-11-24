package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TopUpTransaction represents a top-up transaction record
type TopUpTransaction struct {
	TopUpID       string    `gorm:"column:top_up_id;type:char(36);primaryKey"`
	BalanceID     string    `gorm:"column:balance_id;type:char(36);not null"`
	Amount        int       `gorm:"column:amount;type:int;not null"`
	BalanceBefore int       `gorm:"column:balance_before;type:int;not null"`
	BalanceAfter  int       `gorm:"column:balance_after;type:int;not null"`
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

// BeforeCreate hook to generate a UUID for new users
func (t *TopUpTransaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.TopUpID = uuid.New().String()
	return
}
