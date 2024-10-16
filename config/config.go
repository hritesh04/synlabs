package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn        string
	Secret     string
}

func SetupEnv() (cfg AppConfig, err error) {

	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	httpPort := ":" + os.Getenv("PORT")

	if len(httpPort) < 2 {
		log.Println("HTTP_PORT not found using default port :3000")
		httpPort = ":3000"
	}

	Dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("USER_NAME"), os.Getenv("PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	if len(Dsn) < 1 {
		return AppConfig{}, errors.New("DSN variables not found")
	}

	appSecret := os.Getenv("SECRET")
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("app secret not found")
	}

	return AppConfig{ServerPort: httpPort, Dsn: Dsn, Secret: appSecret}, nil
}
