package usecase

import (
	"errors"
	"ewallet/internal/repository"
	"ewallet/internal/repository/model"
	"ewallet/internal/usecase/entity"
	"ewallet/pkg/connection"
	"ewallet/pkg/jwt"
	"time"
)

type UserUsecase interface {
	RegisterUser(firstName, lastName, phoneNumber, address, pin string) (*entity.User, error)
	LoginUser(username, pin string) (*entity.Token, error)
}

type userUseCaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUseCaseImpl{
		userRepository: userRepository,
	}
}

// RegisterUser registers a new user in the system
// RegisterUser handles the registration of a new user
func (s userUseCaseImpl) RegisterUser(firstName, lastName, phoneNumber, address, pin string) (*entity.User, error) {
	// Check if phone number already exists
	conn := connection.New()

	var existingUser model.User
	if err := conn.DB().Where("phone_number = ?", phoneNumber).First(&existingUser).Error; err == nil {
		return nil, errors.New("phone number already registered")
	}

	// Create the new user
	userModel := model.User{
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
		Address:     address,
		Pin:         pin,
		CreatedAt:   time.Now(),
		Balance:     &model.UserBalance{Balance: 0},
	}
	if err := s.userRepository.CreateUser(&userModel); err != nil {
		return nil, err
	}

	user := entity.NewUser(
		userModel.UserID,
		userModel.FirstName,
		userModel.LastName,
		userModel.PhoneNumber,
		userModel.Address,
		userModel.CreatedAt,
	)

	return &user, nil
}

// LoginUser checks credentials and returns user if valid
func (s userUseCaseImpl) LoginUser(username, pin string) (*entity.Token, error) {

	userModel, err := s.userRepository.LoginUser(username, pin)
	if err != nil {
		return nil, err
	}

	// // Check password
	// err = bcrypt.CompareHashAndPassword([]byte(userModel.Pin), []byte(pin))
	// if err != nil {
	// 	return nil, err
	// }

	// get jwt token
	accessToken, err := jwt.GenerateAccessToken(userModel.UserID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.GenerateRefreshToken(userModel.UserID)
	if err != nil {
		return nil, err
	}

	token := entity.NewToken(accessToken, refreshToken)
	return &token, nil
}
