package com.example.producer;

import com.rabbitmq.client.*;
import java.nio.charset.StandardCharsets;

public class Producer {
    private static final String EXCHANGE_NAME = "events-exchange";

    public static void main(String[] argv) throws Exception {
        ConnectionFactory factory = new ConnectionFactory();
        factory.setHost("rabbitmq");
        factory.setUsername("admin");
        factory.setPassword("admin");

        try (Connection connection = factory.newConnection();
             Channel channel = connection.createChannel()) {

            // Declarar el exchange fanout y durable
            channel.exchangeDeclare(EXCHANGE_NAME, BuiltinExchangeType.FANOUT, true);

            for (int i = 0; i < 5; i++) {
                String message = String.format("{\"source\":\"producer-java\",\"message\":\"event %d\"}", i);
                channel.basicPublish(EXCHANGE_NAME, "", 
                    MessageProperties.PERSISTENT_TEXT_PLAIN,
                    message.getBytes(StandardCharsets.UTF_8));
                System.out.println("Sent: " + message);
                Thread.sleep(1000);
            }
        }
    }
}