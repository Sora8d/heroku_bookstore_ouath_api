package postgresql

import (
	"context"
	"fmt"
	"os"

	pgx "github.com/jackc/pgx/v4"
)

type client struct {
	Conn *pgx.Conn
}

var current_client client

func init() {
	var err error
	current_client.Conn, err = pgx.Connect(context.Background(), fmt.Sprintf("postgress://%s:%s@%s/%s",
		os.Getenv("oauth_postgres_username"),
		os.Getenv("oauth_postgres_password"),
		os.Getenv("oauth_postgres_host"),
		os.Getenv("oauth_postgres_schema")))
	if err != nil {
		panic(err)
	}
}

func GetSession() *client {
	return &current_client
}
