package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hritesh04/synlabs/api/rest"
	"github.com/hritesh04/synlabs/internal/ports"
	"github.com/hritesh04/synlabs/internal/repository"
	"github.com/hritesh04/synlabs/internal/services"
)

type adminHandler struct {
	svc ports.AdminService
}

func SetupAdminHandler(rh rest.RestHandler) {

	adminRepo := repository.NewAdminRepository(rh.DB)

	handler := adminHandler{
		svc: services.NewAdminService(adminRepo, rh.Auth),
	}

	adminRoute := rh.Router.Group("/admin")
	adminRoute.Use(rh.Auth.AdminAuth())

	adminRoute.POST("/job", handler.CreateJob)
	adminRoute.GET("/job/:jobID", handler.GetJob)
	adminRoute.GET("/applicants", handler.GetAllUsers)
	adminRoute.GET("/applicant/:applicantID", handler.GetApplicant)
}

func (h *adminHandler) CreateJob(ctx *gin.Context) {

}

func (h *adminHandler) GetJob(ctx *gin.Context) {

}

func (h *adminHandler) GetAllUsers(ctx *gin.Context) {

}

func (h *adminHandler) GetApplicant(ctx *gin.Context) {

}
