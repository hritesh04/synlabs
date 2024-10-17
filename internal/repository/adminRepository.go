package repository

import (
	"github.com/hritesh04/synlabs/internal/ports"
	"go.mongodb.org/mongo-driver/mongo"
)

type adminRepository struct {
	DB *mongo.Database
}

func NewAdminRepository(db *mongo.Database) ports.AdminRepository {
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
