package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hritesh04/synlabs/api/rest"
	"github.com/hritesh04/synlabs/internal/dto"
	"github.com/hritesh04/synlabs/internal/helper"
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
		svc: services.NewUserService(userRepo, rh.Auth, rh.Parser),
	}

	userRoute := rh.Router

	userRoute.POST("/login", handler.Login)
	userRoute.POST("/signup", handler.Signup)
	userRoute.Use(rh.Auth.Authorize())
	userRoute.POST("/uploadResume", handler.UploadResume)
	protectedUserRoute := userRoute.Group("/jobs")
	protectedUserRoute.GET("/", handler.AllJobs)
	protectedUserRoute.GET("/apply", handler.ApplyToJob)
}

func (h *userHandler) Login(ctx *gin.Context) {
	var user dto.LoginRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		helper.ReturnFailed(ctx, http.StatusInternalServerError, fmt.Errorf("error parsing body : %w", err))
		return
	}
	token, err := h.svc.LogIn(user)
	if err != nil {
		helper.ReturnFailed(ctx, http.StatusInternalServerError, fmt.Errorf("login failed : %w", err))
		return
	}
	helper.ReturnSuccess(ctx, http.StatusOK, token)
}

func (h *userHandler) Signup(ctx *gin.Context) {
	var user dto.SignUpRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		helper.ReturnFailed(ctx, http.StatusBadRequest, fmt.Errorf("error parsing body : %w", err))
		return
	}
	if err := h.svc.SignUp(user); err != nil {
		helper.ReturnFailed(ctx, http.StatusInternalServerError, fmt.Errorf("signup failed : %w", err))
		return
	}
	helper.ReturnSuccess(ctx, http.StatusOK, "user created sucessfully")
}

func (h *userHandler) UploadResume(ctx *gin.Context) {
	file, err := ctx.FormFile("resume")
	if err != nil {
		helper.ReturnFailed(ctx, http.StatusBadRequest, err)
		return
	}
	fileType := file.Header.Get("Content-Type")
	if fileType != "application/pdf" && fileType != "application/vnd.openxmlformats-officedocument.wordprocessingml.document" {
		helper.ReturnFailed(ctx, http.StatusBadRequest, "file type not supported")
		return
	}
	userID := ctx.Request.Header.Get("userID")
	if err := h.svc.UploadResume(file, userID); err != nil {
		helper.ReturnFailed(ctx, http.StatusInternalServerError, err)
		return
	}
	helper.ReturnSuccess(ctx, http.StatusOK, "resume uploaded sucessfully")
}

func (h *userHandler) AllJobs(ctx *gin.Context) {
	result, err := h.svc.GetAllJobs()
	if err != nil {
		helper.ReturnFailed(ctx, http.StatusInternalServerError, err)
		return
	}
	helper.ReturnSuccess(ctx, http.StatusOK, result)
}

func (h *userHandler) ApplyToJob(ctx *gin.Context) {
	userID := ctx.GetHeader("userID")
	jobsID := ctx.Query("job_id")
	if err := h.svc.ApplyToJob(userID, jobsID); err != nil {
		helper.ReturnFailed(ctx, http.StatusInternalServerError, err)
		return
	}
	helper.ReturnSuccess(ctx, http.StatusOK, "applied to job sucessfully")
}
