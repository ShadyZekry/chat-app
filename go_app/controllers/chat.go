package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/ShadyZekry/chat-app/services"
	"github.com/labstack/echo/v4"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Chat struct {
	Number int    `json:"number,omitempty"`
	Name   string `json:"name,omitempty"`
}

func CreateChat(c echo.Context) error {
	token := c.Param("token")
	if !services.ValidateToken(token) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid token"})
	}

	chat := new(Chat)
	chat.Name = c.FormValue("name")

	chat.Number = services.Increment(token)

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
