package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@rabbitmq:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs, _ := ch.Consume("events", "", true, false, false, false, nil)

	fmt.Println(" [*] Waiting for messages...")
	for d := range msgs {
		fmt.Printf("Received: %s\n", d.Body)
	}
}
