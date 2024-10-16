package repository

import (
	"github.com/hritesh04/synlabs/internal/ports"
	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) ports.AdminRepository {
	return &adminRepository{
		DB: db,
	}
}

func (r *adminRepository) CreateJob() {

}

func (r *adminRepository) GetJobByID() {

}

func (r *adminRepository) GetAllUsers() {

}

func (r *adminRepository) GetProfileByUserID() {

}
