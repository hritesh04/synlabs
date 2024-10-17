package ports

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hritesh04/synlabs/internal/domain"
	"github.com/hritesh04/synlabs/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService interface {
	Authorize() gin.HandlerFunc
	Validate(string) (jwt.MapClaims, error)
	AdminAuth() gin.HandlerFunc
	GenerateToken(uint, domain.Role) (string, error)
	HashPassword(string) (string, error)
	ComparePassword(string, string) bool
}

type UserService interface {
	SignUp(dto.SignUpRequest) error
	LogIn(dto.LoginRequest) (string, error)
	UploadResume(*multipart.FileHeader, string) error
	GetAllJobs() (*[]domain.Job, error)
	GetJobByID()
	ApplyToJob(string, string) error
}

type AdminService interface {
	CreateJob(*domain.Job) error
	GetJobInfo()
	GetAllUsers()
	GetUserProfile()
}

type UserRepository interface {
	CreateUser(*domain.User) error
	GetUserByEmail(string) (*domain.User, error)
	CheckUserExists(primitive.ObjectID) error
	GetAllJobs() (*[]domain.Job, error)
	AddUserToJob(primitive.ObjectID, primitive.ObjectID) error
	CreateProfile(*domain.Profile) error
}

type AdminRepository interface {
	CreateJob(*domain.Job) error
	GetJobByID()
	GetAllUsers()
	GetProfileByUserID()
}
