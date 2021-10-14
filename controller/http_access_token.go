package controller

import (
	"net/http"

	"github.com/Sora8d/bookstore_utils-go/rest_errors"
	"github.com/Sora8d/heroku_bookstore_oauth_api/domain/access_token"
	at_services "github.com/Sora8d/heroku_bookstore_oauth_api/services/access_token"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetByIdC(*gin.Context)
	CreateC(*gin.Context)
}

type accessTokenHandler struct {
	service at_services.Service
}

func (handler *accessTokenHandler) GetByIdC(c *gin.Context) {
	accesTokenId := c.Param("access_token_id")
	accessToken, err := handler.service.GetById(accesTokenId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) CreateC(c *gin.Context) {
	var atr access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&atr); err != nil {
		restErr := rest_errors.NewBadRequestErr("invalid fields")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, err := handler.service.Create(atr)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func NewHandler(service at_services.Service) AccessTokenHandler {
	return &accessTokenHandler{service: service}
}
