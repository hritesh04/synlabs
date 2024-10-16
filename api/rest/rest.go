package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/hritesh04/synlabs/internal/auth"
	"gorm.io/gorm"
)

type RestHandler struct {
	Router *gin.Engine
	DB     *gorm.DB
	Auth   *auth.AuthService
}
