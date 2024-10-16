package repository

import (
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

func (r *userRepository) CreateUser() {

}

func (r *userRepository) GetUserByID() {

}

func (r *userRepository) GetAllJobs() {

}

func (r *userRepository) AddUserToJob() {

}

func (r *userRepository) CreateProfile() {

}
