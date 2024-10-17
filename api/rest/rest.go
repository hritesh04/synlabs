package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/hritesh04/synlabs/internal/auth"
	"github.com/hritesh04/synlabs/package/parser"
	"go.mongodb.org/mongo-driver/mongo"
)

type RestHandler struct {
	Router *gin.Engine
	DB     *mongo.Database
	Auth   *auth.AuthService
	Parser *parser.ResumeParser
}
