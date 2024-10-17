package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hritesh04/synlabs/api/rest"
	"github.com/hritesh04/synlabs/internal/domain"
	"github.com/hritesh04/synlabs/internal/helper"
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
	var job domain.Job
	if err := ctx.ShouldBindJSON(&job); err != nil {
		helper.ReturnFailed(ctx, http.StatusBadRequest, err)
		return
	}
	if err := h.svc.CreateJob(&job); err != nil {
		helper.ReturnFailed(ctx, http.StatusInternalServerError, err)
		return
	}
	helper.ReturnSuccess(ctx, http.StatusOK, "job created successfully")
}

func (h *adminHandler) GetJob(ctx *gin.Context) {
	jobID := ctx.Param("jobID")
	result, err := h.svc.GetJobInfo(jobID)
	if err != nil {
		helper.ReturnFailed(ctx, http.StatusInternalServerError, err)
		return
	}
	helper.ReturnSuccess(ctx, http.StatusOK, result)
}

func (h *adminHandler) GetAllUsers(ctx *gin.Context) {

}

func (h *adminHandler) GetApplicant(ctx *gin.Context) {

}
