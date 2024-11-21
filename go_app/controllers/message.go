package controllers

import (
	"database/sql"
	"net/http"

	"github.com/ShadyZekry/chat-app/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type Message struct {
	Number           int    `json:"number,omitempty"`
	Name             string `json:"name,omitempty"`
	ChatNumber       int    `json:"chat_number,omitempty"`
	ApplicationToken string `json:"application_token,omitempty"`
	CreatedAt        string `json:"created_at,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
}

func CreateMessage(c echo.Context) error {
	message := new(Message)
	message.Name = c.FormValue("name")

	app_token := c.Param("token")
	chat_number := c.Param("chat_number")
	message.Number = services.Increment(app_token, chat_number)

	services.Publish(*message, "message")
	return c.JSON(http.StatusCreated, message)
}

func GetMessage(c echo.Context) error {
	db := services.InitDB()
	row := db.QueryRow(
		"select * from messages where application_token=? AND application_token=? AND number=?",
		c.Param("application_token"),
		c.Param("chat_number"),
		c.Param("message_number"),
	)

	message := new(Message)
	err := row.Scan(&message.Number, &message.Name, &message.ChatNumber, &message.ApplicationToken, &message.CreatedAt, &message.UpdatedAt)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Message not found"})
	} else {
		failOnError(err, "Failed to get Message")
	}

	return c.JSON(http.StatusOK, message)
}

func UpdateMessage(c echo.Context) error {
	message := new(Message)
	message.Name = c.FormValue("name")

	services.Publish(*message, "message")
	return c.JSON(http.StatusOK, message)
}
