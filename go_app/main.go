package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Chat struct {
	Number int    `json:"number,omitempty"`
	Name   string `json:"name,omitempty"`
}

type Message struct {
	Number int    `json:"number,omitempty"`
	Name   string `json:"name,omitempty"`
}

var chats []Chat

func createChat(c echo.Context) error {
	chat := new(Chat)
	chat.Name = c.FormValue("name")
	chat.Number = 5

	chats = append(chats, *chat)
	return c.JSON(http.StatusCreated, chats)
}

func main() {
	echo := echo.New()

	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	api := echo.Group("/api/v1")

	api.POST("/applications/:token/chats", createChat)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	echo.Logger.Fatal(echo.Start(":" + httpPort))
}
