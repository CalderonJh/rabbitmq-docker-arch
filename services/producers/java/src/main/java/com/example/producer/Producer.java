package com.example.producer;
import com.rabbitmq.client.*;

public class Producer {
    private final static String EXCHANGE_NAME = "events-topic-exchange";

    public static void main(String[] argv) throws Exception {
        ConnectionFactory factory = new ConnectionFactory();
        factory.setHost("rabbitmq");
        factory.setUsername("admin");
        factory.setPassword("admin");

        try (Connection connection = factory.newConnection();
             Channel channel = connection.createChannel()) {

            channel.exchangeDeclare(EXCHANGE_NAME, BuiltinExchangeType.TOPIC, true);

            int counter = 0;
            while (true) {
                String message = String.format(
                    "{\"source\":\"producer-java\",\"message\":\"event %d\"}", counter++
                );

                channel.basicPublish(
                    EXCHANGE_NAME,
                    "events.java",
                    MessageProperties.PERSISTENT_TEXT_PLAIN,
                    message.getBytes("UTF-8")
                );

                System.out.println("Sent: " + message);
                Thread.sleep(5000);
            }
        }
    }
}
