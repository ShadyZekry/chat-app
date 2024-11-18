package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/ShadyZekry/chat-app/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Chat struct {
	Number int    `json:"number,omitempty"`
	Name   string `json:"name,omitempty"`
}

type Message struct {
	Number int    `json:"number,omitempty"`
	Name   string `json:"name,omitempty"`
}

func createChat(c echo.Context) error {
	chat := new(Chat)
	chat.Name = c.FormValue("name")

	token := c.Param("token")
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
