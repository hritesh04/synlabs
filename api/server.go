package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hritesh04/synlabs/api/rest"
	"github.com/hritesh04/synlabs/api/rest/handlers"
	"github.com/hritesh04/synlabs/config"
	"github.com/hritesh04/synlabs/internal/auth"
	"github.com/hritesh04/synlabs/package/parser"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupServer(cfg config.AppConfig) {

	router := gin.New()
	router.Use(gin.Logger())
	router.MaxMultipartMemory = 5 << 20

	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(cfg.Dsn))
	if err != nil {
		panic(err)
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	db := client.Database("recruitment")
	authService := auth.New(cfg.Secret)
	parser := parser.NewResumeParser(cfg.ApiKey, cfg.ParserUrl)

	rh := rest.RestHandler{
		Router: router,
		DB:     db,
		Auth:   authService,
		Parser: parser,
	}

	setupRoutes(rh)

	router.Run(cfg.ServerPort)
}

func setupRoutes(rh rest.RestHandler) {
	handlers.SetupUserHandler(rh)
	handlers.SetupAdminHandler(rh)
}
