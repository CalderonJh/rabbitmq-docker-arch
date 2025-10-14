package main

import (
	"encoding/json"
	"fmt"
	"log"

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

	q, err := ch.QueueDeclare("events-go", true, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(q.Name, "", "events-exchange", false, nil)
	failOnError(err, "Failed to bind queue")

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	fmt.Println(" [*] Waiting for messages...")
	for d := range msgs {
		var event Event
		json.Unmarshal(d.Body, &event)
		fmt.Printf("Received from %s: %s\n", event.Source, event.Message)
		d.Ack(false)
	}
}
