package repository

import (
	"errors"
	"ewallet/internal/repository/model"
	"ewallet/pkg/connection"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	LoginUser(phoneNumber string, pin string) (*model.User, error)
	FindByID(userID string) (*model.User, error)
}

type userRepositorySql struct {
}

func NewUserRepository() UserRepository {
	return &userRepositorySql{}
}

// CreateUser saves a new user to the database
func (r userRepositorySql) CreateUser(user *model.User) error {
	conn := connection.New()
	if err := conn.DB().Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("phone number already registered")
		}
		return err
	}
	return nil
}

// LoginUser retrieves a user from the database based on phone number and PIN
func (r userRepositorySql) LoginUser(phoneNumber string, pin string) (*model.User, error) {
	conn := connection.New()
	var user model.User
	if err := conn.DB().Where("phone_number = ? AND pin = ?", phoneNumber, pin).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid phone number or PIN")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepositorySql) FindByID(userID string) (*model.User, error) {
	conn := connection.New()
	var user model.User
	if err := conn.DB().First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
