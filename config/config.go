package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn        string
	Secret     string
	ParserUrl  string
	ApiKey     string
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

	Dsn := os.Getenv("DSN")

	if len(Dsn) < 1 {
		return AppConfig{}, errors.New("DSN variables not found")
	}

	apiKey := os.Getenv("API_KEY")
	if len(apiKey) < 1 {
		return AppConfig{}, errors.New("api key not found")
	}
	parserUrl := os.Getenv("PARSER_URL")
	if len(parserUrl) < 1 {
		return AppConfig{}, errors.New("parser url not found")
	}
	appSecret := os.Getenv("SECRET")
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("app secret not found")
	}

	return AppConfig{ServerPort: httpPort, Dsn: Dsn, Secret: appSecret, ParserUrl: parserUrl, ApiKey: apiKey}, nil
}
