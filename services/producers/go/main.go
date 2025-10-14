package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// Connection parameters
const (
	rabbitURL    = "amqp://admin:admin@rabbitmq:5672/"
	exchangeName = "events-topic-exchange"
	routingKey   = "events.go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func connectRabbit() (*amqp.Connection, *amqp.Channel) {
	var conn *amqp.Connection
	var ch *amqp.Channel
	var err error

	// Retry loop until RabbitMQ is available
	for {
		conn, err = amqp.Dial(rabbitURL)
		if err == nil {
			ch, err = conn.Channel()
			if err == nil {
				break
			}
		}
		log.Println("Waiting for RabbitMQ...")
		time.Sleep(3 * time.Second)
	}
	return conn, ch
}

func main() {
	for {
		conn, ch := connectRabbit()
		defer conn.Close()
		defer ch.Close()

		err := ch.ExchangeDeclare(
			exchangeName, // name
			"topic",      // type
			true,         // durable
			false,        // auto-delete
			false,        // internal
			false,        // no-wait
			nil,          // arguments
		)
		failOnError(err, "Failed to declare exchange")

		counter := 0
		for {
			msg := map[string]string{
				"source":  "producer-go",
				"message": fmt.Sprintf("event %d", counter),
			}
			body, _ := json.Marshal(msg)

			err = ch.Publish(
				exchangeName, // exchange
				routingKey,   // routing key
				false,        // mandatory
				false,        // immediate
				amqp.Publishing{
					ContentType:  "application/json",
					Body:         body,
					DeliveryMode: amqp.Persistent,
				},
			)
			if err != nil {
				log.Println("Publish failed, reconnecting...")
				conn.Close()
				break // Exit inner loop → reconnect
			}

			log.Printf("Sent: %s\n", string(body))
			counter++
			time.Sleep(5 * time.Second)
		}
	}
}
