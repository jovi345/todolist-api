package user

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/todos-api/jovi345/token"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CheckEmailAvailabilty(email string) bool
	RegisterUser(user UserRegistrationInput) (User, error)
	Login(input UserLoginInput) (User, error)
	RefreshToken(refreshToken string) (string, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CheckEmailAvailabilty(email string) bool {
	user, _ := s.repository.FindByEmail(email)

	return user.Email != ""
}

func (s *service) RegisterUser(input UserRegistrationInput) (User, error) {
	existedUser := s.CheckEmailAvailabilty(input.Email)
	if existedUser {
		return User{}, errors.New("email is not available")
	}

	if input.Password != input.ConfirmPassword {
		return User{}, errors.New("password do not match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return User{}, err
	}

	id := uuid.New().String()
	createdAt := time.Now()
	updatedAt := createdAt

	user := User{
		ID:        id,
		Username:  input.Username,
		Email:     input.Email,
		Password:  string(hashedPassword),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	newUser, err := s.repository.Save(user)
	if err != nil {
		return User{}, err
	}

	return newUser, nil
}

func (s *service) Login(input UserLoginInput) (User, error) {
	foundUser, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return User{}, errors.New("wrong email or password 1")
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(input.Password))
	log.Print(input.Password)
	if err != nil {
		return User{}, errors.New("wrong email or password 2")
	}

	refreshToken, err := token.GenerateRefreshToken(input.Email)
	if err != nil {
		return User{}, err
	}

	data := UserUpdatedData{
		RefreshToken: refreshToken,
	}

	user, err := s.repository.UpdateUser(foundUser.Email, data)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *service) RefreshToken(refreshToken string) (string, error) {
	user, err := s.repository.FindByRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	accessToken, err := token.GenerateAccessToken(user.Email)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
