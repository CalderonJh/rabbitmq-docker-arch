const amqp = require("amqplib");

(async () => {
  const RABBIT_URL = "amqp://admin:admin@rabbitmq:5672";
  const connection = await amqp.connect(RABBIT_URL);
  const channel = await connection.createChannel();

  await channel.assertExchange("events-topic-exchange", "topic", { durable: true });
  await channel.assertQueue("events-node-2", { durable: true });
  await channel.bindQueue("events-node-2", "events-topic-exchange", "events.node2");

  console.log(" [*] Waiting for Node events...");
  channel.consume("events-node-2", (msg) => {
    if (msg) {
      const data = JSON.parse(msg.content.toString());
      console.log(`[Node consumer 2] Received from ${data.source}: ${data.message}`);
      channel.ack(msg);
    }
  });
})();
