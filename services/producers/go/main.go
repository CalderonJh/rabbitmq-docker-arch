package main

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@rabbitmq:5672/")
	failOnError(err)
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err)
	defer ch.Close()

	q, _ := ch.QueueDeclare("events", true, false, false, false, nil)

	for i := 0; i < 5; i++ {
		body, _ := json.Marshal(map[string]string{"source": "producer-go", "message": fmt.Sprintf("event %d", i)})
		ch.Publish("", q.Name, false, false, amqp.Publishing{ContentType: "application/json", Body: body})
		fmt.Println("Sent:", string(body))
	}
}
