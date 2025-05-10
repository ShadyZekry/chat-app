package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/ShadyZekry/chat-app/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type Chat struct {
	Number           int    `json:"number,omitempty"`
	Name             string `json:"name,omitempty"`
	ApplicationToken string `json:"application_token,omitempty"`
	CreatedAt        string `json:"created_at,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
	MessagesCount    int    `json:"messages_count,omitempty"`
}

func CreateChat(c echo.Context) error {
	chat := new(Chat)
	chat.Name = c.FormValue("name")

	token := c.Param("token")
	chat.Number = services.Increment(token, "chats")
	services.Set(token, strconv.Itoa(chat.Number), "0")
	chat.ApplicationToken = token

	services.Publish(*chat, "chats")
	return c.JSON(http.StatusCreated, chat)
}

func GetChat(c echo.Context) error {
	db := services.InitDB()
	row := db.QueryRow("select * from chats where application_token=? AND number=?", c.Param("token"), c.Param("chat_number"))

	chat := new(Chat)
	err := row.Scan(&chat.Number, &chat.Name, &chat.ApplicationToken, &chat.CreatedAt, &chat.UpdatedAt, &chat.MessagesCount)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Chat not found"})
	} else {
		failOnError(err, "Failed to get chat")
	}

	return c.JSON(http.StatusOK, chat)
}

func UpdateChat(c echo.Context) error {
	chat := new(Chat)
	chat.Name = c.FormValue("name")

	services.Publish(*chat, "chats")
	return c.JSON(http.StatusCreated, chat)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
