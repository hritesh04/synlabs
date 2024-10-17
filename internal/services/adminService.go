package services

import (
	"time"

	"github.com/hritesh04/synlabs/internal/domain"
	"github.com/hritesh04/synlabs/internal/ports"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type adminService struct {
	Repo ports.AdminRepository
	Auth ports.AuthService
}

func NewAdminService(repo ports.AdminRepository, auth ports.AuthService) *adminService {
	return &adminService{
		Repo: repo,
		Auth: auth,
	}
}

func (s *adminService) CreateJob(data *domain.Job) error {
	data.Applicants = []primitive.ObjectID{}
	data.PostedOn = time.Now()
	if err := s.Repo.CreateJob(data); err != nil {
		return err
	}
	return nil
}

func (s *adminService) GetJobInfo(jobID string) (*domain.Job, error) {
	jobObjID, err := primitive.ObjectIDFromHex(jobID)
	if err != nil {
		return nil, err
	}
	result, err := s.Repo.GetJobByID(jobObjID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *adminService) GetAllUsers() {

}

func (s *adminService) GetUserProfile() {

}
