package services

import "github.com/hritesh04/synlabs/internal/ports"

type userService struct {
	Repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *userService {
	return &userService{
		Repo: repo,
	}
}

func (s *userService) SignUp() {

}

func (s *userService) LogIn() {

}

func (s *userService) UploadResume() {

}

func (s *userService) GetAllJobs() {

}

func (s *userService) GetJobByID() {

}
