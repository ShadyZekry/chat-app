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
	api.GET("/applications/:token/chats/:chat_number", controllers.GetChat, middlewares.ValidateApplication)
	api.PUT("/applications/:token/chats/:chat_number", controllers.UpdateChat, middlewares.ValidateApplication, middlewares.ValidateChat)

	api.POST("/applications/:token/chats/:chat_number/messages",
		controllers.CreateMessage,
		middlewares.ValidateApplication, middlewares.ValidateChat,
	)
	api.GET(
		"/applications/:token/chats/:chat_number/messages/:message_number",
		controllers.GetMessage,
		middlewares.ValidateApplication, middlewares.ValidateChat,
	)
	api.PUT(
		"/applications/:token/chats/:chat_number/messages/:message_number",
		controllers.UpdateMessage,
		middlewares.ValidateApplication, middlewares.ValidateChat, middlewares.ValidateMessage,
	)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	echo.Logger.Fatal(echo.Start(":" + httpPort))
}
