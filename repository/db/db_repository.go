package db

import (
	"context"

	"github.com/Sora8d/bookstore_utils-go/rest_errors"
	"github.com/Sora8d/heroku_bookstore_oauth_api/clients/postgresql"
	"github.com/Sora8d/heroku_bookstore_oauth_api/domain/access_token"
	pgx "github.com/jackc/pgx/v4"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=$1;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES ($1, $2, $3, $4);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=$1 WHERE access_token=$2;"
)

type DbRepository interface {
	GetById(access_token.AccessToken) (*access_token.AccessToken, rest_errors.RestErr)
	Create(access_token.AccessToken) rest_errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) rest_errors.RestErr
}

type dbRepository struct {
}

func (dbr *dbRepository) GetById(at access_token.AccessToken) (*access_token.AccessToken, rest_errors.RestErr) {
	// TODO: implement get accesstoken from PostGresSQL
	session := postgresql.GetSession().Conn
	if err := session.QueryRow(context.Background(), queryGetAccessToken, at.AccessToken).Scan(
		&at.AccessToken,
		&at.UserId,
		&at.ClientId,
		&at.Expires,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, rest_errors.NewBadRequestErr("Access Token Not Found")
		}
		return nil, rest_errors.NewInternalServerError("Error getting AT id", err)
	}
	return &at, nil
}

func (dbr *dbRepository) Create(at access_token.AccessToken) rest_errors.RestErr {
	session := postgresql.GetSession().Conn
	if _, err := session.Exec(context.Background(), queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	); err != nil {
		return rest_errors.NewInternalServerError("Error generating access token", err)
	}
	return nil
}

func (dbr *dbRepository) UpdateExpirationTime(at access_token.AccessToken) rest_errors.RestErr {
	session := postgresql.GetSession().Conn
	if _, err := session.Exec(context.Background(), queryUpdateExpires,
		at.Expires,
		at.AccessToken,
	); err != nil {
		return rest_errors.NewInternalServerError("There was an error updating the exp time", err)
	}
	return nil
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
