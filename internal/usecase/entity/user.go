package entity

import (
	"time"
)

type User struct {
	UserID      string    `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_date"`
}

func NewUser(
	userID string,
	firstName string,
	lastName string,
	phoneNumber string,
	address string,
	createdAt time.Time,
) User {
	return User{
		UserID:      userID,
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
		Address:     address,
		CreatedAt:   createdAt,
	}
}
