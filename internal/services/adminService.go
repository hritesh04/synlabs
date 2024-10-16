package services

import "github.com/hritesh04/synlabs/internal/ports"

type adminService struct {
	Repo ports.AdminRepository
}

func NewAdminService(repo ports.AdminRepository) *adminService {
	return &adminService{
		Repo: repo,
	}
}

func (s *adminService) CreateJob() {

}

func (s *adminService) GetJobInfo() {

}

func (s *adminService) GetAllUsers() {

}

func (s *adminService) GetUserProfile() {

}
