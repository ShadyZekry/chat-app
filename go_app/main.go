package main

import (
	"os"

	"github.com/ShadyZekry/chat-app/controllers"
	"github.com/ShadyZekry/chat-app/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	echo := echo.New()

	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	api := echo.Group("/api/v1")

	api.POST("/applications/:token/chats", controllers.CreateChat, middlewares.ValidateApplication)
	api.GET("/applications/:token/chats/:number", controllers.GetChat, middlewares.ValidateApplication)
	api.PUT("/applications/:token/chats/:number", controllers.UpdateChat, middlewares.ValidateApplication, middlewares.ValidateChat)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	echo.Logger.Fatal(echo.Start(":" + httpPort))
}
