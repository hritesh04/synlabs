package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hritesh04/synlabs/api/rest"
	"github.com/hritesh04/synlabs/api/rest/handlers"
	"github.com/hritesh04/synlabs/config"
	"github.com/hritesh04/synlabs/internal/auth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupServer(cfg config.AppConfig) {

	router := gin.Default()
	router.Use(gin.Logger())

	db, err := gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(); err != nil {
		log.Println("database migration failed")
	}

	authService := auth.New(cfg.Secret)

	rh := rest.RestHandler{
		Router: router,
		DB:     db,
		Auth:   authService,
	}

	setupRoutes(rh)

	router.Run(cfg.ServerPort)
}

func setupRoutes(rh rest.RestHandler) {
	handlers.SetupUserHandler(rh)
	handlers.SetupAdminHandler(rh)
}
