package rest

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/Sora8d/bookstore_utils-go/rest_errors"
	"github.com/Sora8d/heroku_bookstore_oauth_api/config"
	"github.com/Sora8d/heroku_bookstore_oauth_api/domain/users"
	rest "github.com/go-resty/resty/v2"
)

var (
	usersRestClient *rest.Client = rest.New()
	userServerUrl   string       = config.Config["UsersURI"]
)

func NewRepository() UsersRepository {
	return &usersRepository{}
}

type UsersRepository interface {
	LoginUser(string, string) (*users.User, rest_errors.RestErr)
}

type usersRepository struct{}

func (rR *usersRepository) LoginUser(email string, password string) (*users.User, rest_errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	usersRestClient.SetTimeout(5 * time.Second)
	usersRestClient.SetHostURL(userServerUrl)
	restreq := usersRestClient.R()
	restreq.SetHeader("Content-Type", "application/json")
	restreq.Method = rest.MethodPost
	restreq.Body = request
	restreq.URL = "/users/login"

	response, err := restreq.Send()
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error in the restclient functionality", err)
	}
	if response == nil || response.Body() == nil {
		return nil, rest_errors.NewInternalServerError("invalid restclient response when trying to login user", errors.New("nil response"))
	}
	if response.IsError() {
		var restErr rest_errors.RestErr
		err := json.Unmarshal(response.Body(), &restErr)
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to log into user", err)
		}
		return nil, restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Body(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal users response", errors.New("users server response invalid"))
	}
	return &user, nil
}
