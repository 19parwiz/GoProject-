package service

import (
	"bookstore/internal/models"
	"bookstore/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// UserService Struct
type UserService struct {
	UserRepo *repository.UserRepository
}

// Constructor Function: NewUserService
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

// RegisterUser Method  user registeration with hashing password
func (s *UserService) RegisterUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user.Password = string(hashedPassword)
	return s.UserRepo.CreateUser(user)
}

// Login (password verification)         LoginUser Method
func (s *UserService) LoginUser(email, password string) (*models.User, error) {
	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// Password Verification
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

// Новый метод: Получить пользователя по ID
func (s *UserService) GetUserByID(userID int) (*models.User, error) {
	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
