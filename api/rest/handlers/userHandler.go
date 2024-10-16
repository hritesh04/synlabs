package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hritesh04/synlabs/api/rest"
	"github.com/hritesh04/synlabs/internal/ports"
	"github.com/hritesh04/synlabs/internal/repository"
	"github.com/hritesh04/synlabs/internal/services"
)

type userHandler struct {
	svc ports.UserService
}

func SetupUserHandler(rh rest.RestHandler) {
	userRepo := repository.NewUserRepository(rh.DB)

	handler := userHandler{
		svc: services.NewUserService(userRepo, rh.Auth),
	}

	userRoute := rh.Router

	userRoute.POST("/login", handler.Login)
	userRoute.POST("/signup", handler.Signup)
	userRoute.POST("/uploadResune", handler.UploadResume)
	protectedUserRoute := userRoute.Group("/jobs")
	protectedUserRoute.Use(rh.Auth.Authorize())
	protectedUserRoute.GET("/", handler.AllJobs)
	protectedUserRoute.GET("/apply", handler.ApplyToJob)
}

func (h *userHandler) Login(ctx *gin.Context) {

}

func (h *userHandler) Signup(ctx *gin.Context) {

}

func (h *userHandler) UploadResume(ctx *gin.Context) {

}

func (h *userHandler) AllJobs(ctx *gin.Context) {

}

func (h *userHandler) ApplyToJob(ctx *gin.Context) {

}
