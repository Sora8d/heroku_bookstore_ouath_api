package config

import (
	"os"
)

type config map[string]string

func init() {

	Config = config{
		"database": os.Getenv("DATABASE_URL"),
		"UsersURI": os.Getenv("UsersURI"),
		"address":  os.Getenv("address"),
		"port":     os.Getenv("PORT"),
	}
}

var Config config
