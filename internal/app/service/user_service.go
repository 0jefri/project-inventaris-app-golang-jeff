package service

import (
	"time"

	"github.com/project-app-inventaris/internal/app/repository"
	"github.com/project-app-inventaris/internal/model"
	"github.com/project-app-inventaris/internal/model/dto"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterNewUser(payload *model.User) (*dto.UserResponse, error)
	FindByUsername(username string) (*model.User, error)
	FindByUsernamePassword(username string, password string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterNewUser(payload *model.User) (*dto.UserResponse, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	password := string(bytes)

	payload.Password = password

	user, err := s.repo.Create(payload)

	userResponse := dto.UserResponse{
		ID:               user.ID,
		Username:         user.Username,
		Email:            user.Email,
		Password:         user.Password,
		RegistrationDate: user.RegistrationDate,
		LastLogin:        time.Now(),
	}

	return &userResponse, err
}

func (s *userService) FindByUsername(username string) (*model.User, error) {
	return s.repo.GetUsername(username)
}

func (s *userService) FindByUsernamePassword(username string, password string) (*model.User, error) {
	return s.repo.GetUsernamePassword(username, password)
}
