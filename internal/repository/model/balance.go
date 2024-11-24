package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserBalance struct {
	ID        string    `gorm:"column:id;type:char(36);primaryKey"`
	UserID    string    `gorm:"column:user_id;type:char(36);not null"`
	Balance   int       `gorm:"column:balance;type:int;default:0"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	User      *User     `gorm:"foreignKey:UserID"`
}

// BeforeCreate hook to generate a UUID for new users
func (u *UserBalance) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
