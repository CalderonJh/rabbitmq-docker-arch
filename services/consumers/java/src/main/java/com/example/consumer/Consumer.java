package com.example.consumer;

import com.rabbitmq.client.*;

public class Consumer {
    private final static String EXCHANGE_NAME = "events-topic-exchange";

    public static void main(String[] argv) throws Exception {
        ConnectionFactory factory = new ConnectionFactory();
        factory.setHost("rabbitmq");
        factory.setUsername("admin");
        factory.setPassword("admin");
        Connection connection = factory.newConnection();
        Channel channel = connection.createChannel();

        channel.exchangeDeclare(EXCHANGE_NAME, BuiltinExchangeType.TOPIC, true);
        channel.queueDeclare("events-java", true, false, false, null);
        channel.queueBind("events-java", EXCHANGE_NAME, "events.java");

        System.out.println(" [*] Waiting for Java events...");
        DeliverCallback deliverCallback = (consumerTag, delivery) -> {
            String message = new String(delivery.getBody(), "UTF-8");
            System.out.println("[Java consumer] Received: " + message);
            channel.basicAck(delivery.getEnvelope().getDeliveryTag(), false);
        };
        channel.basicConsume("events-java", false, deliverCallback, consumerTag -> {});
    }
}