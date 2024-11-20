package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ShadyZekry/chat-app/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	amqp "github.com/rabbitmq/amqp091-go"
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

	publishChat(*chat)
	return c.JSON(http.StatusCreated, chat)
}

func GetChat(c echo.Context) error {
	db := services.InitDB()
	row := db.QueryRow("select * from chats where application_token=? AND number=?", c.Param("token"), c.Param("number"))

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

	publishChat(*chat)
	return c.JSON(http.StatusCreated, chat)
}

func publishChat(chat Chat) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"chats", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body, err := json.Marshal(chat)
	failOnError(err, "Failed to marshal chat")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
