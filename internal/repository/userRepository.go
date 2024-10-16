package repository

import (
	"github.com/hritesh04/synlabs/internal/domain"
	"github.com/hritesh04/synlabs/internal/ports"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) CreateUser(data *domain.User) error {
	result := r.DB.Create(data)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := r.DB.First(&user, "email = ?", email)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAllJobs() {

}

func (r *userRepository) AddUserToJob() {

}

func (r *userRepository) CreateProfile() {

}
