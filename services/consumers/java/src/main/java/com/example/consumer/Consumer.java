package com.example.consumer;

import com.rabbitmq.client.*;
import java.nio.charset.StandardCharsets;

public class Consumer {
    private static final String EXCHANGE_NAME = "events-exchange";
    private static final String QUEUE_NAME = "events-java";

    public static void main(String[] argv) throws Exception {
        ConnectionFactory factory = new ConnectionFactory();
        factory.setHost("rabbitmq");
        factory.setUsername("admin");
        factory.setPassword("admin");

        Connection connection = factory.newConnection();
        Channel channel = connection.createChannel();

        // Declarar el exchange y la cola durables
        channel.exchangeDeclare(EXCHANGE_NAME, BuiltinExchangeType.FANOUT, true);
        channel.queueDeclare(QUEUE_NAME, true, false, false, null);
        channel.queueBind(QUEUE_NAME, EXCHANGE_NAME, "");

        System.out.println(" [*] Waiting for messages...");

        DeliverCallback deliverCallback = (consumerTag, delivery) -> {
            String message = new String(delivery.getBody(), StandardCharsets.UTF_8);
            System.out.println("Received: " + message);
            channel.basicAck(delivery.getEnvelope().getDeliveryTag(), false);
        };

        channel.basicConsume(QUEUE_NAME, false, deliverCallback, consumerTag -> {});
    }
}
