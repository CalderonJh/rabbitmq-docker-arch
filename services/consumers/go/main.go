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

	// Asegura que la cola exista con las mismas propiedades
	_, err = ch.QueueDeclare(
		"events", // nombre
		true,     // durable
		false,    // autoDelete
		false,    // exclusive
		false,    // noWait
		nil,      // args
	)
	if err != nil {
		panic(err)
	}

	msgs, err := ch.Consume("events", "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(" [*] Waiting for messages...")
	for d := range msgs {
		fmt.Printf("Received: %s\n", d.Body)
	}
}
