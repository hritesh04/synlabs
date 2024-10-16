package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hritesh04/synlabs/internal/dto"
)

type AuthService interface {
	Authorize() gin.HandlerFunc
	Validate(string) (jwt.MapClaims, error)
	AdminAuth() gin.HandlerFunc
	GenerateToken()
	HashPassword(string) (string, error)
	ComparePassword(string, string) bool
}

type UserService interface {
	SignUp(dto.SignUpRequest)
	LogIn()
	UploadResume()
	GetAllJobs()
	GetJobByID()
}

type AdminService interface {
	CreateJob()
	GetJobInfo()
	GetAllUsers()
	GetUserProfile()
}

type UserRepository interface {
	CreateUser()
	GetUserByID()
	GetAllJobs()
	AddUserToJob()
	CreateProfile()
}

type AdminRepository interface {
	CreateJob()
	GetJobByID()
	GetAllUsers()
	GetProfileByUserID()
}
