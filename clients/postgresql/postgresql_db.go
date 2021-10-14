package postgresql

import (
	"context"
	"fmt"

	"github.com/Sora8d/bookstore_utils-go/logger"
	"github.com/Sora8d/heroku_bookstore_oauth_api/config"
	pgx "github.com/jackc/pgx/v4"
)

type client struct {
	Conn *pgx.Conn
}

var current_client client

func init() {
	var err error
	current_client.Conn, err = pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s/%s",
		config.Config["oauth_postgres_username"],
		config.Config["oauth_postgres_password"],
		config.Config["oauth_postgres_host"],
		config.Config["oauth_postgres_schema"]))
	if err != nil {
		logger.Error("Error connecting to the database, shutting down the app", err)
		panic(err)
	}
}

func GetSession() *client {
	return &current_client
}
