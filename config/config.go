package config

import (
	"os"
)

type config map[string]string

func init() {

	Config = config{
		"database": os.Getenv("database"),
		"UsersURI": os.Getenv("UsersURI"),
		"address":  os.Getenv("address"),
	}
}

var Config config
