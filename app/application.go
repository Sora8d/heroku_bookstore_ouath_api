package app

import (
	"github.com/Sora8d/heroku_bookstore_oauth_api/config"
	"github.com/Sora8d/heroku_bookstore_oauth_api/controller"
	"github.com/Sora8d/heroku_bookstore_oauth_api/repository/db"
	"github.com/Sora8d/heroku_bookstore_oauth_api/repository/rest"
	at_services "github.com/Sora8d/heroku_bookstore_oauth_api/services/access_token"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	atService := at_services.NewService(db.NewRepository(), rest.NewRepository())
	atHandler := controller.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByIdC)
	router.POST("/oauth/access_token", atHandler.CreateC)
	router.Run(config.Config["address"])
}
