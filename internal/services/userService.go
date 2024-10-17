package services

import (
	"github.com/hritesh04/synlabs/internal/domain"
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

func (s *userService) SignUp(data dto.SignUpRequest) error {
	user := &domain.User{
		Name:            data.Name,
		Email:           data.Email,
		UserType:        domain.Role(data.UserType),
		ProfileHeadline: data.ProfileHeadline,
		Address:         data.Address,
	}

	hash, err := s.Auth.HashPassword(data.Password)
	if err != nil {
		return err
	}
	user.PasswordHash = hash

	if err := s.Repo.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func (s *userService) LogIn(data dto.LoginRequest) (string, error) {
	user, err := s.Repo.GetUserByEmail(data.Email)
	if err != nil {
		return "", err
	}
	isPasswordCorrect := s.Auth.ComparePassword(user.PasswordHash, data.Password)

	if !isPasswordCorrect {
		return "", nil
	}

	token, err := s.Auth.GenerateToken(user.ID, user.UserType)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) UploadResume() {

}

func (s *userService) GetAllJobs() (*[]domain.Job, error) {
	jobs, err := s.Repo.GetAllJobs()
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (s *userService) GetJobByID() {

}

func (s *userService) ApplyToJob(userID, jobID string) error {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid user_id")
	}
	jobObjID, err := primitive.ObjectIDFromHex(jobID)
	if err != nil {
		return fmt.Errorf("invalid job_id")
	}

	if err := s.Repo.CheckUserExists(userObjID); err != nil {
		return fmt.Errorf("user not found")
	}

	if err := s.Repo.AddUserToJob(userObjID, jobObjID); err != nil {
		return err
	}
	return nil
}
