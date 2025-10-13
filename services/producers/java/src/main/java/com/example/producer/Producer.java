package com.example.producer;

import com.rabbitmq.client.*;

public class Producer {
  private static final String QUEUE_NAME = "events";

  public static void main(String[] argv) throws Exception {
    ConnectionFactory factory = new ConnectionFactory();
    factory.setHost("rabbitmq");
    factory.setUsername("admin");
    factory.setPassword("admin");

    try (Connection connection = factory.newConnection();
        Channel channel = connection.createChannel()) {
      channel.queueDeclare(QUEUE_NAME, false, false, false, null);
      for (int i = 0; i < 5; i++) {
        String message = "event " + i;
        channel.basicPublish("", QUEUE_NAME, null, message.getBytes());
        System.out.println("Sent: " + message);
      }
    }
  }
}
