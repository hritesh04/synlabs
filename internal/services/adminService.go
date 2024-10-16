package services

import "github.com/hritesh04/synlabs/internal/ports"

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

func (s *adminService) CreateJob() {

}

func (s *adminService) GetJobInfo() {

}

func (s *adminService) GetAllUsers() {

}

func (s *adminService) GetUserProfile() {

}
