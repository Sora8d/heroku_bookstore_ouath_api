package db

import (
	"context"
	"errors"

	"github.com/Sora8d/bookstore_utils-go/logger"
	"github.com/Sora8d/bookstore_utils-go/rest_errors"
	"github.com/Sora8d/heroku_bookstore_oauth_api/clients/postgresql"
	"github.com/Sora8d/heroku_bookstore_oauth_api/domain/access_token"
	pgx "github.com/jackc/pgx/v4"
)

const (
	queryGetAccessToken      = "SELECT access_token, user_id, client_id, permissions, expires FROM access_tokens WHERE access_token=$1;"
	queryCreateAccessToken   = "INSERT INTO access_tokens(access_token, user_id, client_id, permissions, expires) VALUES ($1, $2, $3, $4, $5);"
	queryUpdateExpires       = "UPDATE access_tokens SET expires=$1 WHERE access_token=$2;"
	queryAuthenticateSecrets = "SELECT client_id, permissions FROM not_so_secret WHERE client_id=$1 AND secret=$2"
)

type DbRepository interface {
	GetById(access_token.AccessToken) (*access_token.AccessToken, rest_errors.RestErr)
	Create(access_token.AccessToken) rest_errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) rest_errors.RestErr
	AuthSecret(int64, string) (interface{}, rest_errors.RestErr)
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
		&at.Admin,
		&at.Expires,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, rest_errors.NewBadRequestErr("Access Token Not Found")
		}
		logger.Error("Error in the GetById function inside db_repository.go,", err)
		return nil, rest_errors.NewInternalServerError("Error getting AT id", errors.New("dabase error"))
	}
	return &at, nil
}

func (dbr *dbRepository) Create(at access_token.AccessToken) rest_errors.RestErr {
	session := postgresql.GetSession().Conn
	if _, err := session.Exec(context.Background(), queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Admin,
		at.Expires,
	); err != nil {
		return rest_errors.NewBadRequestErr("Error generating access token")
	}
	return nil
}

func (dbr *dbRepository) UpdateExpirationTime(at access_token.AccessToken) rest_errors.RestErr {
	session := postgresql.GetSession().Conn
	if _, err := session.Exec(context.Background(), queryUpdateExpires,
		at.Expires,
		at.AccessToken,
	); err != nil {
		logger.Error("Error in the UpdateExpirationTime function inside db_repository.go,", err)
		return rest_errors.NewInternalServerError("There was an error updating the exp time", err)
	}
	return nil
}

func (dbr *dbRepository) AuthSecret(client_id int64, secret string) (interface{}, rest_errors.RestErr) {
	session := postgresql.GetSession().Conn
	row := session.QueryRow(context.Background(), queryAuthenticateSecrets, client_id, secret)
	var result bool
	if err := row.Scan(&result); err != nil {
		return nil, rest_errors.NewBadRequestErr("Bad Credentials")
	}
	return result, nil
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
