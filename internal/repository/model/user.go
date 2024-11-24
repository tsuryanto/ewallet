package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID      string         `gorm:"column:user_id;type:char(36);primaryKey"`
	FirstName   string         `gorm:"column:first_name;type:varchar(255);not null"`
	LastName    string         `gorm:"column:last_name;type:varchar(255);not null"`
	PhoneNumber string         `gorm:"column:phone_number;type:varchar(15);unique;not null"`
	Address     string         `gorm:"column:address;type:text"`
	Pin         string         `gorm:"column:pin;type:varchar(255);not null"` // Hashed
	CreatedAt   time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;index"`
	Balance     *UserBalance   `gorm:"foreignKey:UserID"`
}

// BeforeCreate hook to generate a UUID for new users
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UserID = uuid.New().String()
	return
}
