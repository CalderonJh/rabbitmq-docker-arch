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

    ch.ExchangeDeclare("events-topic-exchange", "topic", true, false, false, false, nil)
    q, err := ch.QueueDeclare("events-go", true, false, false, false, nil)
    failOnError(err)
    ch.QueueBind(q.Name, "events.go", "events-topic-exchange", false, nil)

    msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
    failOnError(err)

    fmt.Println(" [*] Waiting for Go events...")
    for d := range msgs {
        var data map[string]string
        json.Unmarshal(d.Body, &data)
        fmt.Printf("[Go consumer] Received from %s: %s\n", data["source"], data["message"])
        d.Ack(false)
    }
}
