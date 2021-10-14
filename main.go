package main

import (
	"github.com/Sora8d/heroku_bookstore_oauth_api/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("test_envs.env")
	app.StartApplication()
}
