package rest

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/Sora8d/bookstore_utils-go/rest_errors"
	"github.com/Sora8d/heroku_bookstore_oauth_api/domain/users"

	mock "github.com/jarcoal/httpmock"
)

type mockjson struct{}

type mockError struct {
	Message int    `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

type mockUser struct {
	Id        string
	FirstName string
	LastName  string
}

func TestMain(m *testing.M) {
	mock.ActivateNonDefault(usersRestClient.GetClient())
	log.Println("about to start test cases...")
	os.Exit(m.Run())
}

/*
Couldnt test time out yet corretly
func TestLoginUserTimeoutFromApi(t *testing.T) {
	mock.Reset()
	resp, err := mock.NewJsonResponder(-1, mockjson{})
	if err != nil {
		fmt.Println(err)
		t.Error("Incorrect Test")
		return
	}
	mock.RegisterResponder(http.MethodPost, "https://api.bookstore.com/users/login", resp)
	repository := usersRepository{}

	user, resterr := repository.LoginUser("email@gmail.com", "password")

	if !(user == nil && resterr != nil && http.StatusInternalServerError != resterr.Status && resterr.Message != "invalid restclient response when trying to login user") {
		fmt.Println("user: ", user)
		fmt.Println("err: ", resterr)
		fmt.Println("err status: ", resterr.Status)
		fmt.Println("err message: ", resterr.Message)
		t.Error("Incorrect Response")
	}
}
*/
func TestLoginUserInvalidErrorInterface(t *testing.T) {
	mock.Reset()
	resp, err := mock.NewJsonResponder(400, mockError{Message: 123, Status: 400, Error: "yheaa"})
	if err != nil {
		fmt.Println(err)
		t.Error("Incorrect Test")
		return
	}
	mock.RegisterResponder(http.MethodPost, "https://api.bookstore.com/users/login", resp)
	repository := usersRepository{}

	user, resterr := repository.LoginUser("email@gmail.com", "password")

	if !(user == nil && resterr != nil && resterr.Status() == http.StatusInternalServerError && resterr.Message() == "invalid error interface when trying to log into user") {
		fmt.Println("user:", user)
		fmt.Println("err:", resterr)
		t.Error("Incorrect Response")
	}

}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	mock.Reset()
	resp, err := mock.NewJsonResponder(400, rest_errors.NewBadRequestErr("Incorrect user credentials"))
	if err != nil {
		fmt.Println(err)
		t.Error("Incorrect Test")
		return
	}
	mock.RegisterResponder(http.MethodPost, "https://api.bookstore.com/users/login", resp)
	repository := usersRepository{}

	user, resterr := repository.LoginUser("email@gmail.com", "password")
	if !(user == nil && resterr != nil && resterr.Status() == http.StatusBadRequest && resterr.Message() == "Incorrect user credentials") {
		fmt.Println("user: ", user)
		fmt.Println("err: ", resterr)
		t.Error("Incorrect Response")
	}
	return
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	mock.Reset()
	resp, err := mock.NewJsonResponder(201, mockUser{
		Id:        "4",
		FirstName: "Jhon",
		LastName:  "Salmon",
	})
	if err != nil {
		fmt.Println(err)
		t.Error("Incorrect Test")
		return
	}
	mock.RegisterResponder(http.MethodPost, "https://api.bookstore.com/users/login", resp)
	repository := usersRepository{}

	user, resterr := repository.LoginUser("email@gmail.com", "password")
	if !(user == nil && resterr != nil && resterr.Status() == http.StatusInternalServerError && resterr.Message() == "error when trying to unmarshal users response") {
		fmt.Println("user: ", user)
		fmt.Println("err: ", resterr)
		t.Error("Incorrect Response")
	}
	return
}

func TestLoginUserNoError(t *testing.T) {
	mock.Reset()
	resp, err := mock.NewJsonResponder(201, users.User{
		Id:        4,
		FirstName: "Jhon",
		LastName:  "Salmon",
		Email:     "Jhonahood",
	})
	if err != nil {
		fmt.Println(err)
		t.Error("Incorrect Test")
		return
	}
	mock.RegisterResponder(http.MethodPost, "https://api.bookstore.com/users/login", resp)
	repository := usersRepository{}

	user, resterr := repository.LoginUser("email@gmail.com", "password")
	if user == nil && resterr != nil {
		fmt.Println("user: ", user)
		fmt.Println("err: ", resterr)
		t.Error("Incorrect Response")
	}
	return
}
