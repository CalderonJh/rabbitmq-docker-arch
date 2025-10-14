package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type Event struct {
	Source  string `json:"source"`
	Message string `json:"message"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare("events-exchange", "fanout", true, false, false, false, nil)
	failOnError(err, "Failed to declare an exchange")

	for i := 0; i < 5; i++ {
		msg := Event{Source: "producer-go", Message: fmt.Sprintf("event %d", i)}
		body, _ := json.Marshal(msg)
		err = ch.Publish(
			"events-exchange", "", false, false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
				DeliveryMode: amqp.Persistent, // mensaje persistente
			})
		failOnError(err, "Failed to publish a message")
		fmt.Println("Sent:", string(body))
		time.Sleep(1 * time.Second)
	}
}
