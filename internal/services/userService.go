package services

import (
	"github.com/hritesh04/synlabs/internal/dto"
	"github.com/hritesh04/synlabs/internal/ports"
)

type userService struct {
	Repo ports.UserRepository
	Auth ports.AuthService
}

func NewUserService(repo ports.UserRepository, auth ports.AuthService) *userService {
	return &userService{
		Repo: repo,
		Auth: auth,
	}
}

func (s *userService) SignUp(data dto.SignUpRequest) {

}

func (s *userService) LogIn() {

}

func (s *userService) UploadResume() {

}

func (s *userService) GetAllJobs() {

}

func (s *userService) GetJobByID() {

}
