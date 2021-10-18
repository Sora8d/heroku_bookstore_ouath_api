package config

import (
	"os"

	"github.com/Sora8d/bookstore_utils-go/logger"
	"github.com/joho/godotenv"
)

type config map[string]string

func init() {
	if err := godotenv.Load("test_envs.env"); err != nil {
		logger.Error("Error loading environment variables", err)
		panic(err)
	}

	Config = config{
		"oauth_postgres_username": os.Getenv("oauth_postgres_username"),
		"oauth_postgres_password": os.Getenv("oauth_postgres_password"),
		"oauth_postgres_schema":   os.Getenv("oauth_postgres_schema"),
		"oauth_postgres_host":     os.Getenv("oauth_postgres_host"),
		"UsersURI":                os.Getenv("UsersURI"),
		"address":                 os.Getenv("address"),
	}
}

var Config config
