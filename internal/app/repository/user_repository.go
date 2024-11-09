package repository

import (
	"fmt"

	"github.com/project-app-inventaris/internal/model"
	"github.com/project-app-inventaris/utils/exception"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	// BaseRepository[model.User]
	// BaseRepositoryPaging[model.User]
	Create(payload *model.User) (*model.User, error)
	GetUsernamePassword(username, password string) (*model.User, error)
	GetUsername(username string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(payload *model.User) (*model.User, error) {
	user := model.User{
		ID:       payload.ID,
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
	}

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err

	}
	fmt.Println(user)

	return &user, nil
}

func (r *userRepository) GetUsernamePassword(username, password string) (*model.User, error) {

	user, err := r.GetUsername(username)

	if err != nil {
		return nil, exception.ErrNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, exception.ErrInvalidUsernamePassword
	}

	return user, err
}

func (r *userRepository) GetUsername(username string) (*model.User, error) {
	user := model.User{}

	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
